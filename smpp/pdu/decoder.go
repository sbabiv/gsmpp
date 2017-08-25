package pdu

import (
	"encoding/binary"
	"errors"
	"net"
	"fmt"
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

/*func Decode(c net.Conn) error {
	h, err := HeaderDecoder(c)
	if err != nil {
		return err
	}
}*/

func Decode(b []byte) {
	h := &Header{
		Length:   binary.BigEndian.Uint32(b[0:4]),
		Id:       binary.BigEndian.Uint32(b[4:8]),
		Status:   binary.BigEndian.Uint32(b[8:12]),
		Sequence: binary.BigEndian.Uint32(b[12:16]),
	}

	n, systemId := decodeCoctet(b[16:], 16)
	fmt.Printf("n: v%, systemid: v%", n, systemId)
	fmt.Printf("len %v", h.Length)
}

func decodeCoctet(b []byte, offset int) (int, string) {
	for i, v := range b {
		if v == 0 {
			return offset + i + 1, string(b[:i+1])
		}
	}
	return 0, ""
}
