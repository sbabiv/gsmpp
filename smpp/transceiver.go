package smpp

import (
	"time"
	"strconv"
	"github.com/gsmpp/smpp/pdu"
	"github.com/gsmpp/smpp/tcp"
	"github.com/gsmpp/smpp/decoder"
	"fmt"
)


type Transceiver struct {
	host string
	port int
	systemId string
	password string
	systemType string
	interfaceVersion byte
	addrTon byte
	addrNpi byte
	addressRange string
	conn *tcp.Client

	ChannelState chan Event
	enquirelinkTicker *time.Ticker
}

func NewTransceiver(host string, port int, systemId string, password string, systemType string) *Transceiver {
	return &Transceiver{
		host:             host,
		port:             port,
		systemId:         systemId,
		password:         password,
		systemType:       systemType,
		interfaceVersion: pdu.SMPP_INTERFACE_VERSION,
		ChannelState:     make(chan Event, 1),
		conn:             &tcp.Client{},
	}
}

func (t *Transceiver) Bind() error {
	err := t.conn.Dial("tcp", t.host + ":" + strconv.Itoa(t.port))
	if err != nil {
		t.ChannelState <- NewEvent(CONN_FAIL)
		return err
	}

	t.ChannelState <- NewEvent(CONNECTED)
	bind := pdu.NewBindTransceiverCommand(t.systemId, t.password, t.systemType, t.addressRange, t.addrTon, t.addrNpi)

	_, err = t.conn.Write(bind.Bytes())
	if err != nil {
		t.ChannelState <- NewEvent(BIND_FAIL)
		t.conn.Close()
		return err
	}

	go t.reader()

	return nil
}

func (t *Transceiver)Unbind() {
	t.conn.Write(pdu.NewUnbindCommand().Bytes())
}

func (t *Transceiver) sendEnquireLink()  {
	for range t.enquirelinkTicker.C {
		n, err := t.conn.Write(pdu.NewEnquireLinkCommand().Bytes())
		if err != nil || n == 0 {
			t.ChannelState <- NewEvent(SEND_ENQUIRE_LINK_ERR)
			t.conn.Close()

			return
		}
		t.ChannelState <- NewEvent(SEND_ENQUIRE_LINK)
	}
}

func (t *Transceiver) reader() {
	for {
		h, err := decoder.HeaderDecoder(t.conn)
		if err != nil {
			t.ChannelState <- NewEvent(READ_PDU_ERR)
		}
		fmt.Printf("len: %v, id: %v, seq: %v\n", h.Length, h.Id, h.Sequence)

		switch h.Id {
		case pdu.BIND_TRANSCEIVER_RESP:
			val, err := decoder.BindTransceiverDecoder(t.conn, int(h.GetBodyLen()))
			if err != nil {
				t.ChannelState <- NewEvent(BIND_FAIL)
			}
			fmt.Printf("bind resp. system id: %v\n", val.SystemId)
			t.ChannelState <- NewEvent(BOUND_TRX)

			//t.enquirelinkTicker = time.NewTicker(time.Second * 30)
			//go t.sendEnquireLink()

		case pdu.ENQUIRE_LINK_RESP:
			t.ChannelState <- NewEvent(SEND_ENQUIRE_LINK_ERR)
		case pdu.ENQUIRE_LINK:
			t.conn.Write(pdu.NewEnquireLinkRespCommand(h.Sequence).Bytes())
		default:
			if h.Sequence == 13 {
				fmt.Println("")
			}

			_, err := decoder.Skip(t.conn, int(h.GetBodyLen()))
			if err != nil {
				t.ChannelState <- NewEvent(READ_PDU_ERR)
			}else {
				t.ChannelState <- NewEvent(SKIP_PDU)
			}
		}
	}
}
