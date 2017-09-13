package pdu

import (
	"bytes"

)

const (
	BIND_TRANSCEIVER uint32 = 0x00000009
	BIND_TRANSCEIVER_RESP = 0x80000009
	SMPP_INTERFACE_VERSION byte = 0x34
	ENQUIRE_LINK uint32 = 0x00000015
	ENQUIRE_LINK_RESP uint32 = 0x80000015

	/*tlv tag*/
	SC_INTERFACE_VERSION = 0x0210
)

type BindTransceiverCommand struct {
	header *Header
	body   []byte
}

func NewBindTransceiverCommand(
	systemId, password, systemType, addressRange string,
	addrTon, addrNpi byte) *BindTransceiverCommand {

	var b bytes.Buffer
	b.WriteString(coctet(systemId))
	b.WriteString(coctet(password))
	b.WriteString(coctet(systemType))
	b.WriteByte(SMPP_INTERFACE_VERSION)
	b.WriteByte(addrTon)
	b.WriteByte(addrNpi)
	b.WriteString(coctet(addressRange))
	h := NewHeader(uint32(b.Len()), BIND_TRANSCEIVER, 0)

	return &BindTransceiverCommand{header: h, body: b.Bytes()}
}

func (c *BindTransceiverCommand) Bytes() []byte {
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