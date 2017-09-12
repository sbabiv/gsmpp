package pdu

import (
	"bytes"

)

const (
	BindTransceiver uint32 = 0x00000009
	BindTransceiverResp = 0x80000009
	SmppInterfaceVersion byte = 0x34
	enquire_link uint32 = 0x00000015
	enquire_link_resp uint32 = 0x80000015

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
	b.WriteByte(SmppInterfaceVersion)
	b.WriteByte(addrTon)
	b.WriteByte(addrNpi)
	b.WriteString(coctet(addressRange))
	h := NewHeader(uint32(b.Len()), BindTransceiver, 0)

	return &BindTransceiverCommand{header: h, body: b.Bytes()}
}

func (c *BindTransceiverCommand) Bytes() []byte {
	h := c.header.Bytes()
	b := make([]byte, len(h)+len(c.body))
	copy(b, h)
	copy(b[len(h):], c.body)

	return b
}

type Unbind struct {
	Header
}
