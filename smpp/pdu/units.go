package pdu

import (
	"sync/atomic"
	"time"
	"bytes"
	"encoding/binary"
)

const (
	HeaderLength = 16
	NullTerminated = "\x00"
)

var seq uint32 = uint32(time.Now().Unix())

type OptionalParameters map[uint16]*TLV

func coctet(v string) string {
	return v + NullTerminated
}

func sequenceInc() uint32 {
	return atomic.AddUint32(&seq,1)
}

type Header struct {
	Length uint32
	Id uint32
	Status uint32
	Sequence uint32
}

func NewHeader(len uint32, id uint32, status uint32) *Header {
	return &Header{Length: len + HeaderLength, Id: id, Status: status, Sequence: sequenceInc()}
}

func (h *Header)Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, h)

	return buf.Bytes()
}

func (h *Header)GetBodyLen() uint32 {
	return h.Length - HeaderLength
}

type TLV struct {
	Tag    uint16
	Length uint16
	Value  []byte
}

type BindTransceiverResp struct {
	SystemId string
	OptionalParameters
}


type BindCommand struct {
	header *Header
	body   []byte
}

func NewBindTrxCommand(
	systemId, password, systemType, addressRange string,
	addrTon, addrNpi byte) *BindCommand {

	var b bytes.Buffer
	b.WriteString(coctet(systemId))
	b.WriteString(coctet(password))
	b.WriteString(coctet(systemType))
	b.WriteByte(SMPP_INTERFACE_VERSION)
	b.WriteByte(addrTon)
	b.WriteByte(addrNpi)
	b.WriteString(coctet(addressRange))
	h := NewHeader(uint32(b.Len()), BIND_TRANSCEIVER, 0)

	return &BindCommand{header: h, body: b.Bytes()}
}

func (c *BindCommand) Bytes() []byte {
	h := c.header.Bytes()
	b := make([]byte, len(h)+len(c.body))
	copy(b, h)
	copy(b[len(h):], c.body)

	return b
}

type UnbindCommand struct {
	Header
}

func NewUnbindCommand() *UnbindCommand{
	return &UnbindCommand{}
}

type EnquireLinkCommand struct {
	Header
}

func NewEnquireLinkCommand() *EnquireLinkCommand {
	return &EnquireLinkCommand{
		Header{
			Length:   HeaderLength,
			Id:       ENQUIRE_LINK,
			Status:   0,
			Sequence: sequenceInc(),
		},
	}
}

func NewEnquireLinkRespCommand(sequence uint32) *EnquireLinkCommand {
	return &EnquireLinkCommand{
		Header{
			Length:   HeaderLength,
			Id:       ENQUIRE_LINK_RESP,
			Status:   0,
			Sequence: sequence,
		},
	}
}

func (e *EnquireLinkCommand)Bytes() []byte  {
	return e.Header.Bytes()
}