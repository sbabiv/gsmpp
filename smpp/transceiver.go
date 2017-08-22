package smpp

import (
	"net"
	"gsmpp/gsmpp/smpp/pdu"
	"time"
	"fmt"
	"strconv"
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
	}
}

func (t *Transceiver) Bind() {
	conn, err := net.Dial("tcp", t.host + ":" + strconv.Itoa(t.port))
	if err != nil {
		t.ChannelState <- NewEvent(CONN_FAIL, "connection failed")
		return
	}
	_, err = conn.Write([]byte("bind"))
	if err != nil {
		t.ChannelState <- NewEvent(BIND, "binding event")
		conn.Close()
		return
	}
	t.enquirelinkTicker = time.NewTicker(time.Second * 2)
	go t.sendEnquireLink()
	t.ChannelState <- NewEvent(BIND, "bind event")
}

func (t *Transceiver) sendEnquireLink()  {
	for item := range t.enquirelinkTicker.C {
		fmt.Println(item)
	}
}
