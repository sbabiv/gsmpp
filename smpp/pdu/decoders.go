package pdu

/*import (
	"encoding/binary"
	"errors"
	"net"
)*/
/*

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
*/
























