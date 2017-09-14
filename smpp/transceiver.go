package smpp

import (
	"time"
	"strconv"
	"github.com/gsmpp/smpp/pdu"
	"github.com/gsmpp/smpp/decoder"
	"net"
	"github.com/gsmpp/smpp/events"
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
	conn net.Conn

	ChannelState chan events.Event
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
		ChannelState:     make(chan events.Event, 1),
	}
}

func (t *Transceiver) Bind() error {
	var err error
	t.conn, err = net.Dial("tcp", t.host + ":" + strconv.Itoa(t.port))
	if err != nil {
		t.ChannelState <- events.NewEvent(events.CONN_FAIL)
		return err
	}
	t.ChannelState <- events.NewEvent(events.CONNECTED)
	_, err = t.conn.Write(pdu.NewBindTrxCommand(t.systemId, t.password, t.systemType, t.addressRange, t.addrTon, t.addrNpi).Bytes())
	if err != nil {
		t.ChannelState <- events.NewEvent(events.BIND_FAIL)
		t.Close()
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
		_, err := t.conn.Write(pdu.NewEnquireLinkCommand().Bytes())
		if err != nil {
			t.Close()
			return
		}
		t.ChannelState <- events.NewEvent(events.SEND_ENQUIRE_LINK)
	}
}

func (t *Transceiver) Close(){
	if t.enquirelinkTicker != nil {
		t.enquirelinkTicker.Stop()
	}
	if t.conn != nil {
		t.conn.Close()
		t.conn = nil
	}
	t.ChannelState <- events.NewEvent(events.DISCONNECTED)
}

func (t *Transceiver) reader() {
	for {
		h, err := decoder.HeaderDecoder(t.conn)
		if err != nil {
			t.ChannelState <- events.NewEvent(events.READ_PDU_ERR)
			t.Close()
			return
		}

		switch h.Id {

		case pdu.BIND_TRANSCEIVER_RESP:
			_, err := decoder.BindTransceiverDecoder(t.conn, int(h.GetBodyLen()))
			if err != nil {
				t.ChannelState <- events.NewEvent(events.READ_PDU_ERR)
				t.conn.Close()
				return
			}
			t.ChannelState <- events.NewEvent(events.BOUND_TRX)
			t.enquirelinkTicker = time.NewTicker(time.Second * 30)
			go t.sendEnquireLink()

		case pdu.ENQUIRE_LINK_RESP:
			t.ChannelState <- events.NewEvent(events.SEND_ENQUIRE_LINK_RESP)

		case pdu.ENQUIRE_LINK:
			t.conn.Write(pdu.NewEnquireLinkRespCommand(h.Sequence).Bytes())

		default:
			_, err := decoder.Skip(t.conn, int(h.GetBodyLen()))
			if err != nil {
				t.ChannelState <- events.NewEvent(events.READ_PDU_ERR)
			} else {
				t.ChannelState <- events.NewEvent(events.SKIP_PDU)
			}
		}
	}
}
