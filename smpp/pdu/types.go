package pdu

import (
	"encoding/hex"
	"bytes"
)

type Field struct {
	FieldName
	val []byte
}

func F(n FieldName, v interface{}) *Field {
	return &Field{n, encode(v)}
}

func (f *Field) Len() int {
	return len(f.val)
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

func (this *TLV) String() string {
	return string(this.Value)
}
/*func (this *TLV) Byte() string {
	return string(this.Value)
}
func (this *TLV) Short() string {
	return string(this.Value)
}
func (this *TLV) Int() string {
	return string(this.Value)
}*/

type Optionals map[TagName]*TLV

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

type Unit struct {
	*Header
	data []byte
}

func (this *Unit) Bytes() []byte {
	return this.data
}

func (this *Unit) Dump() string {
	return hex.Dump(this.data)
}

func NewUnit(id uint32, status uint32, seq *uint32, fields Fields, optionals Optionals) *Unit {
	var b bytes.Buffer

	b.Write(fields.Encode())
	b.Write(optionals.Encode())
	h := NewHeader(uint32(b.Len()), id, status, seq)

	return &Unit{h, append(h.Bytes(), b.Bytes()...)}
}

type UnitRsp struct {
	*Header
	*Body
}

func (this *UnitRsp) Dump() string {
	var b bytes.Buffer
	b.Write(this.Header.Bytes())
	b.Write(this.Body.Raw)

	return hex.Dump(b.Bytes())
}

func NewBindTrx(systemId, password, systemType, addressRange string, addrTon, addrNpi byte) *Unit {
	return NewUnit(BIND_TRANSCEIVER, 0x0, nil, Fields{
		F(SystemId, systemId),
		F(Password, password),
		F(SystemType, systemType),
		F(InterfaceVersion, SMPP_INTERFACE_VERSION),
		F(AddrTon, addrTon),
		F(AddrNpi, addrNpi),
		F(AddressRange, addressRange),
	}, nil)
}

func NewEnquireLink() *Unit {
	return NewUnit(ENQUIRE_LINK,0,nil,nil,nil)
}

func NewEnquireLinkResp(seq uint32) *Unit {
	//5.1.3 command_status
	return NewUnit(ENQUIRE_LINK_RESP, ESME_ROK, &seq, nil, nil)
}

func NewDeliverSmResp(seq uint32) *Unit {
	//5.1.3 command_status
	return NewUnit(DELIVER_SM_RESP, ESME_ROK, &seq, nil, nil)
}

func NewUnbind() *Unit {
	return NewUnit(UNBIND, 0x0, nil, nil, nil)
}

func NewUnbindResp(seq uint32) *Unit {
	return NewUnit(UNBIND_RESP, ESME_ROK, &seq, nil, nil)
}

func NewSubmitSm(message, number string) *Unit {
	return NewUnit(SUBMIT_SM, 0x0, nil, Fields{
		F(ServiceType, ""),
		F(SourceAddrTon, byte(5)),
		F(SourceAddrNpi, byte(0)),
		F(SourceAddr, ""),
		F(DestAddrNpi, 1),
		F(DestAddrTon, 1),
		F(DestinationAddr, number),
		F(EsmClass, 0x00010011),
		F(ProtocolId, byte(1)),
		F(PriorityFlag, byte(0)),
		F(ScheduleDeliveryTime, ""),
		F(ValidityPeriod, ""),
		F(RegisteredDelivery, 0x00001111),
		F(ReplaceIfPresentFlag, 0),
		F(DataCoding, byte(0x0)),
		F(SmDefaultMsgId, 0),
		F(SmLength, byte(len([]byte(message)))),
		F(ShortMessage, []byte(message)),
	}, nil)
}