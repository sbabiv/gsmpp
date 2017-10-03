package smpp

import (
	"time"
	"strconv"
	"net"
	"github.com/sbabiv/gsmpp/smpp/pdu"
	"github.com/sbabiv/gsmpp/smpp/events"
	"github.com/sbabiv/gsmpp/smpp/pdu/decoders"
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
	t.conn, err = net.Dial("tcp", t.host+":"+strconv.Itoa(t.port))
	if err != nil {
		t.ChannelState <- events.NewEvent(events.CONN_FAIL)
		return err
	}
	t.ChannelState <- events.NewEvent(events.CONNECTED)
	_, err = t.conn.Write(pdu.NewBindTrx(t.systemId, t.password, t.systemType, t.addressRange, t.addrTon, t.addrNpi).Bytes())
	if err != nil {
		t.ChannelState <- events.NewEvent(events.BIND_FAIL)
		t.Close()
		return err
	}

	go t.reader()

	return nil
}

func (t *Transceiver)Unbind() error {
	_, err := t.conn.Write(pdu.NewUnbind().Bytes())
	if err == nil {
		t.enquirelinkTicker.Stop()
	}
	return err
}

func (t *Transceiver) sendEnquireLink()  {
	for range t.enquirelinkTicker.C {
		_, err := t.conn.Write(pdu.NewEnquireLink().Bytes())
		if err != nil {
			t.Close()
			return
		}
		t.ChannelState <- events.NewEvent(events.SEND_ENQUIRE_LINK)
	}
}

func (t *Transceiver) Close() error {
	if t.enquirelinkTicker != nil {
		t.enquirelinkTicker.Stop()
	}
	if t.conn != nil {
		err := t.conn.Close()
		if err != nil {
			return err
		}
		t.conn = nil
	}
	t.ChannelState <- events.NewEvent(events.DISCONNECTED)
	return nil
}

func (t *Transceiver) reader() {
	for {
		h, err := decoders.DecodeHeader(t.conn)
		if err != nil {
			t.ChannelState <- events.NewEvent(events.READ_PDU_ERR)
			t.Close()
			return
		}

		switch h.Id {

		case pdu.BIND_TRANSCEIVER_RESP:
			r, err := decoders.DecodeBindResp(h, t.conn)
			if err != nil {
				t.ChannelState <- events.NewEvent(events.READ_PDU_ERR)
				t.conn.Close()
				return
			}
			fmt.Println(r.Dump())

			t.ChannelState <- events.NewEvent(events.BOUND_TRX)
			t.enquirelinkTicker = time.NewTicker(time.Second * 30)
			go t.sendEnquireLink()


		case pdu.ENQUIRE_LINK:
			t.conn.Write(pdu.NewEnquireLinkResp(h.Sequence).Bytes())
			t.ChannelState <- events.NewEvent(events.RECEIVE_ENQUITE_LINK)

		case pdu.ENQUIRE_LINK_RESP:
			t.ChannelState <- events.NewEvent(events.RECEIVE_ENQUITE_LINK_RESP)

		case pdu.UNBIND:
			t.conn.Write(pdu.NewUnbindResp(h.Sequence).Bytes())
			t.ChannelState <- events.NewEvent(events.UNBIND)
			t.Close()

		case pdu.UNBIND_RESP:
			t.ChannelState <- events.NewEvent(events.UNBIND)
			t.Close()

		case pdu.DELIVER_SM:
			r, err := decoders.DecodeDeliverSm(h, t.conn)
			if err != nil {
				t.ChannelState <- events.NewEvent(events.READ_PDU_ERR)
				t.conn.Close()
				return
			}

			fmt.Println(r.Dump())

			t.ChannelState <- events.NewEvent(events.DELIVER_SM)
			t.conn.Write(pdu.NewDeliverSmResp(h.Sequence).Bytes())

		default:
			decoders.Skip(h, t.conn)
			t.ChannelState <- events.NewEvent(events.SKIP_PDU)
		}
	}
}
