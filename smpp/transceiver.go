package smpp

import "net"


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
	events chan Event
}

func NewTransceiver(host string, port int, systemId string, password string, systemType string) *Transceiver {
	return &Transceiver{host: host, port:port, systemId:systemId, password:password, systemType:systemType, interfaceVersion:34, events:make(chan Event, 1)}
}

func Bind(t *Transceiver) {
	conn, err := net.Dial("tcp", "localhost:8075")
	if err != nil {
		t.events <- Event{Status:1, Message:err.Error()}
		return
	}

	t.conn = conn
	//return Erro
}

func Events(t *Transceiver) chan Event {
	return t.events
}

func Unbind(){

}

/*
func send(t *Transceiver) error {
	t.conn.Write()
}*/
