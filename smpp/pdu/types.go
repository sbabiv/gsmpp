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
	return &Field{n, Encode(v)}
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

type FieldNames []FieldName

func (f FieldNames) Decode([]byte) ([]Field, []TLV) {
	return nil, nil
}

type Fields []*Field

func (this Fields) Encode() []byte {
	var b bytes.Buffer
	for _, f := range this {
		b.Write(f.Bytes())
	}
	return b.Bytes()
}

func Encode(v interface{}) []byte {
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