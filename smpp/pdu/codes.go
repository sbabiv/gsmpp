package pdu

const (
	//commad ids
	BIND_TRANSCEIVER uint32 = 0x00000009
	BIND_TRANSCEIVER_RESP uint32 = 0x80000009

	UNBIND uint32 = 0x00000006
	UNBIND_RESP uint32 = 0x80000006

	ENQUIRE_LINK uint32 = 0x00000015
	ENQUIRE_LINK_RESP uint32 = 0x80000015

	SUBMIT_SM uint32 = 0x00000004
	SUBMIT_SM_RESP uint32 = 0x80000004

	DELIVER_SM uint32 = 0x00000005
	DELIVER_SM_RESP uint32 = 0x80000005

	//


	//5.1.3 command_status
	//no error
	ESME_ROK uint32 = 0x00000000
	//Message Length is invalid
	ESME_RINVMSGLEN uint32 = 0x00000001
	//Command Length is invalid
	ESME_RINVCMDLEN uint32 = 0x00000002
	//Invalid Command ID
	ESME_RINVCMDID uint32 = 0x00000003
	//Incorrect BIND int for given command
	ESME_RINVBNDSTS uint32 = 0x00000004
	//ESME Already in Bound State
	ESME_RALYBND uint32 = 0x00000005
	//Invalid Priority Flad
	ESME_RINVPRTFLG uint32 = 0x00000006
	//Invalid Registred Delivery Flag
	ESME_RINVREGDLVFLG uint32 = 0x00000007
	//System Error
	ESME_RSYSERR uint32 = 0x00000008

	SMPP_INTERFACE_VERSION byte = 0x34
)

type FieldName string

const (
	SystemId                FieldName = "system_id"
	Password                FieldName = "password"
	SystemType              FieldName = "system_type"
	ServiceType             FieldName = "service_type"
	MessageId               FieldName = "message_id"
	InterfaceVersion        FieldName = "interface_version"
	SourceAddrTon           FieldName = "source_addr_ton"
	SourceAddrNpi           FieldName = "source_addr_npi"
	SourceAddr           	FieldName = "source_addr"
	DestAddrTon          	FieldName = "dest_addr_ton"
	DestAddrNpi          	FieldName = "dest_addr_npi"
	DestinationAddr      	FieldName = "destination_addr"
	EsmClass             	FieldName = "esm_class"
	ProtocolId           	FieldName = "protocol_id"
	PriorityFlag         	FieldName = "priority_flag"
	ScheduleDeliveryTime 	FieldName = "schedule_delivery_time"
	ValidityPeriod       	FieldName = "validity_period"
	RegisteredDelivery   	FieldName = "registered_delivery"
	ReplaceIfPresentFlag 	FieldName = "replace_if_present_flag"
	DataCoding           	FieldName = "data_coding"
	SmDefaultMsgId       	FieldName = "sm_default_msg_id"
	SmLength             	FieldName = "sm_length"
	ShortMessage         	FieldName = "short_message"
	AddrTon              	FieldName = "add_ton"
	AddrNpi              	FieldName = "addr_npi"
	AddressRange         	FieldName = "address_range"
)

//Message state const
type State byte

const (
	Enroute State = iota+1
	Delivered
	Expired
	Deleted
	Undeliverable
	Accepted
	Unknown
	Rejected
)

//TagName TLV const
type TagName uint16

const (
	ScInterfaceVersion TagName = 0x0210
	ReceiptedMessageId TagName = 0x001e
	MessageState 	   TagName = 0x0427
)

type Mask byte

const (
	//1011
	EsmDeliv Mask = (2 << 2) | 3
	EsmDelivUdh Mask = (2 << 2) | 3 | (1 << 6)
	//11101
	RegisteredDeliv Mask = (7 << 2) | 1
)

type Multipart byte

const (
	Payload Multipart = 0
	UDH Multipart = 1
	SAR Multipart = 2
)
