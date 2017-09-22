package pdu

import (
	"encoding/binary"
	"errors"
)

type Reader interface {
	Read(b []byte) (n int, err error)
}

func coctetDecoder(b []byte) (int, string) {
	for i, v := range b {
		if v == 0 {
			return i + 1, string(b[:i+1])
		}
	}
	return 0, ""
}

func tlvDecoder(b []byte) (*TLV, error) {
	l := binary.BigEndian.Uint16(b[2:4])
	if len(b)-int(l) < 0 {
		return nil, errors.New("TLV decoding error")
	}

	return &TLV{
		Tag:    binary.BigEndian.Uint16(b[:2]),
		Length: l,
		Value:  b[4:4+l],
	}, nil
}

func HeaderDecoder(r Reader) (*Header, error){
	b := make([]byte, HeaderLength, HeaderLength)
	_, err := r.Read(b)

	if err != nil {
		return nil, err
	}

	return &Header{
		Length:   binary.BigEndian.Uint32(b[0:4]),
		Id:       binary.BigEndian.Uint32(b[4:8]),
		Status:   binary.BigEndian.Uint32(b[8:12]),
		Sequence: binary.BigEndian.Uint32(b[12:16]),
	}, nil
}

func BindTransceiverDecoder(r Reader, len int) (*BindTransceiverResp, error) {
	b := make([]byte, len, len)
	_, err := r.Read(b)
	if err != nil {
		return nil, err
	}

	n, systemId := coctetDecoder(b)
	if n == 0 {
		return nil, errors.New("filed parse bind transceiver resp")
	}

	p := make(OptionalParameters)
	tlv, err := tlvDecoder(b[n:])
	if err != nil {
		return nil, err
	}
	p[tlv.Tag] = tlv

	return &BindTransceiverResp{
		systemId,
		p,
	}, nil
}

func Skip(r Reader, len int) ([]byte, error) {
	if len == 0 {
		return nil, nil
	}
	b := make([]byte, len, len)
	_, err := r.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}






















