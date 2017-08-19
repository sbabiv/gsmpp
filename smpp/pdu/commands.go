package pdu

type BindTransceiverCommand struct {
	Header
	systemId string
	password string
	systemType string
	interfaceVersion byte
	addrTon byte
	addrNpi byte
	addressRange string
}



const (
	BindRransceiverCommand uint32 = 0x00000009
	BindTransceiverResp = 0x80000009
	SmppInterfaceVersion byte = 0x34
)

