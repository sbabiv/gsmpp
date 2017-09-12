package smpp

import (
	"time"
	"fmt"
	"strconv"
	"github.com/gsmpp/smpp/pdu"
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
	conn *tcpClient

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
		interfaceVersion: pdu.SmppInterfaceVersion,
		ChannelState:     make(chan Event, 1),
		conn:             &tcpClient{},
	}
}

func (t *Transceiver) Bind() {
	err := t.conn.Dial("tcp", t.host + ":" + strconv.Itoa(t.port))
	if err != nil {
		t.ChannelState <- NewEvent(CONN_FAIL, "connection failed")
		return
	}

	bind := pdu.NewBindTransceiverCommand(t.systemId, t.password, t.systemType, t.addressRange, t.addrTon, t.addrNpi)

	_, err = t.conn.Write(bind.Bytes())
	if err != nil {
		t.ChannelState <- NewEvent(BIND, "binding event")
		t.conn.Close()
		return
	}
	t.enquirelinkTicker = time.NewTicker(time.Second * 2)
	go t.sendEnquireLink()
	t.ChannelState <- NewEvent(BIND, "bind event")
}

func Unbind() {

}

func (t *Transceiver) sendEnquireLink()  {
	for item := range t.enquirelinkTicker.C {
		fmt.Println(item)
	}
}
