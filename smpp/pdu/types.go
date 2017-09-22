package pdu

import (
	"encoding/hex"
	"bytes"
)

type FieldName string

const (
	SystemId FieldName = "system_id"
	Password FieldName = "password"
	SystemType FieldName = "system_type"
	InterfaceVersion FieldName = "interface_version"
	AddrTon FieldName = "add_ton"
	AddrNpi FieldName = "addr_npi"
	AddressRange = "address_range"
)

type Field struct {
	FieldName
	val []byte
}

func F(n FieldName, v interface{}) *Field {
	return &Field{n, encode(v)}
}

func (f *Field) Bytes() []byte {
	return f.val
}

func (f *Field) String() string {
	return string(f.val)
}

func (f *Field) Dump() string {
	return hex.Dump(f.val)
}

type TLV struct {
	Tag    uint16
	Length uint16
	Value  []byte
}

type Optionals map[FieldName]*TLV

func (this Optionals) Encode() []byte {
	return nil
}

type Fields []*Field

func (this Fields) Encode() []byte {
	var b bytes.Buffer
	for _, f := range this {
		b.Write(f.Bytes())
	}
	return b.Bytes()
}

func encode(v interface{}) []byte {
	switch v.(type) {
	case []byte:
		return v.([]byte)
	case byte:
		return []byte{v.(byte)}
	case string:
		return []byte(coctet(v.(string)))
	default:
		return []byte{0}
	}
}

type FieldNames []FieldName

func (f FieldNames) Decode([]byte) ([]Field, []TLV) {
	return nil, nil
}

type Command struct {
	*Header
	data []byte
}

func (this *Command) Bytes() []byte {
	return this.data
}

func (this *Command) Dump() string {
	return hex.Dump(this.data)
}

func NewCommand(id uint32, status uint32, fields Fields, optionals Optionals) *Command {
	var b bytes.Buffer

	b.Write(fields.Encode())
	b.Write(optionals.Encode())
	h := NewHeader(uint32(b.Len()), id, status)

	return &Command{h, append(h.Bytes(), b.Bytes()...)}
}

func NewBindTrx(systemId, password, systemType, addressRange string, addrTon, addrNpi byte) *Command {
	return NewCommand(BIND_TRANSCEIVER, 0, Fields{
		F(SystemId, systemId),
		F(Password, password),
		F(SystemType, systemType),
		F(InterfaceVersion, SMPP_INTERFACE_VERSION),
		F(AddrTon, addrTon),
		F(AddrNpi, addrNpi),
		F(AddressRange, addressRange),
	}, nil)
}