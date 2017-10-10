package pdu

import (
	"bytes"
	"encoding/binary"
	"net"
	"errors"
)

type FieldNames []FieldName

func (f FieldNames) Decode(buffer *bytes.Buffer) Fields {
	result := make(Fields, len(f), len(f))
	for i, name := range f {
		switch name {
		case
			SystemId,
			SystemType,
			ServiceType,
			SourceAddr,
			AddressRange,
			DestinationAddr,
			ScheduleDeliveryTime,
			ValidityPeriod,
			MessageId:

			b, _ := buffer.ReadBytes(0x00)
			result[i] = F(name, b)

		case
			SourceAddrTon,
			SourceAddrNpi,
			DestAddrTon,
			DestAddrNpi,
			EsmClass,
			ProtocolId,
			PriorityFlag,
			RegisteredDelivery,
			ReplaceIfPresentFlag,
			DataCoding,
			SmDefaultMsgId,
			SmLength:

			b, _ := buffer.ReadByte()
			result[i] = F(name, []byte{b})

		case ShortMessage:
			result[i] = F(name, buffer.Next(int(result[i-1].val[0])))
		}
	}
	return result
}

type Body struct {
	Fields
	Optionals
	Raw []byte
}

type Decoder struct {
	names FieldNames
	buff *bytes.Buffer
}

func NewDecoder(buff *bytes.Buffer, names FieldNames) *Decoder {
	return &Decoder{names, buff}
}

func (this *Decoder) Decode() *Body {
	b := this.buff.Bytes()
	return &Body{
		this.names.Decode(this.buff),
		this.decodeOptionals(),
		b,
	}
}

func (this *Decoder)decodeOptionals() Optionals {
	result := make(Optionals)
	for {
		if this.buff.Len() == 0 {
			break
		}
		tag := binary.BigEndian.Uint16(this.buff.Next(2))
		len := binary.BigEndian.Uint16(this.buff.Next(2))
		val := this.buff.Next(int(len))

		result[TagName(tag)] = &TLV{tag, len, val}
	}
	return result
}

func D(h *Header, conn net.Conn, names FieldNames) (*UnitRsp, error) {
	b := make([]byte, h.GetBodyLen(), h.GetBodyLen())
	n, err := conn.Read(b)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		return nil, errors.New("Error transport protocol")
	}
	buff := bytes.NewBuffer(b)
	d := NewDecoder(buff, names)

	return &UnitRsp{h, d.Decode()}, nil
}




