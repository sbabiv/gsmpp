package pdu

import (
	"encoding/binary"
	"errors"
	"net"
	"fmt"
)

func coctedDecoder(b []byte, offset int) (int, string) {
	for i, v := range b {
		if v == 0 {
			return offset + i + 1, string(b[:i+1])
		}
	}
	return 0, ""
}

func tlvDecoder(b []byte) (*TLV, error) {
	l := binary.BigEndian.Uint16(b[2:4])
	if len(b)-int(l) < 0 {
		return nil, fmt.Errorf("TLV decoding error")
	}

	return &TLV{
		Tag:    binary.BigEndian.Uint16(b[:2]),
		Length: l,
		Value:  b[4:4+l],
	}, nil
}

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

func Decode(b []byte) {
	h := &Header{
		Length:   binary.BigEndian.Uint32(b[0:4]),
		Id:       binary.BigEndian.Uint32(b[4:8]),
		Status:   binary.BigEndian.Uint32(b[8:12]),
		Sequence: binary.BigEndian.Uint32(b[12:16]),
	}

	resp, _ := bindTransceiverRespDecoder(b)
	fmt.Printf("system id: %v, h.len: %v", resp.SystemId, h.Length)
}

func bindTransceiverRespDecoder(b []byte) (*BIND_TRANSCEIVER_RESP, error) {
	n, systemId := coctedDecoder(b[16:], 16)
	if n == 0 {
		return nil, fmt.Errorf("filed parse bind transceiver resp")
	}

	p := make(OptionalParameters)
	tlv, err := tlvDecoder(b[n:])
	if err != nil {
		return nil, err
	}
	p[tlv.Tag] = tlv

	return &BIND_TRANSCEIVER_RESP{
		systemId,
		p,
	}, nil
}


