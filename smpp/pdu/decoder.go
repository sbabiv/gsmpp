package pdu

import (
	"encoding/binary"
	"errors"
	"net"
)

func HeaderDecoder(c net.Conn) (*Header, error) {
	b := make([]byte, HeaderLength)
	n, err := c.Read(b)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, errors.New("Malformed packed")
	}

	return &Header{
		Length:   binary.BigEndian.Uint32(b[0:4]),
		Id:       binary.BigEndian.Uint32(b[4:8]),
		Status:   binary.BigEndian.Uint32(b[8:12]),
		Sequence: binary.BigEndian.Uint32(b[12:16]),
	}, nil
}
