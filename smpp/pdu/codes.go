package pdu


const (
	//commad ids
	BIND_TRANSCEIVER uint32 = 0x00000009
	BIND_TRANSCEIVER_RESP uint32 = 0x80000009

	UNBIND uint32 = 0x00000006
	UNBIND_RESP uint32 = 0x80000006

	ENQUIRE_LINK uint32 = 0x00000015
	ENQUIRE_LINK_RESP uint32 = 0x80000015

	//command status
	//no error
	ESME_ROK int = 0x00000000
	//Message Length is invalid
	ESME_RINVMSGLEN int = 0x00000001
	//Command Length is invalid
	ESME_RINVCMDLEN int = 0x00000002
	//Invalid Command ID
	ESME_RINVCMDID int = 0x00000003
	//Incorrect BIND int for given command
	ESME_RINVBNDSTS int = 0x00000004
	//ESME Already in Bound State
	ESME_RALYBND int = 0x00000005
	//Invalid Priority Flad
	ESME_RINVPRTFLG int = 0x00000006
	//Invalid Registred Delivery Flag
	ESME_RINVREGDLVFLG = 0x00000007
	//System Error
	ESME_RSYSERR int = 0x00000008

	SMPP_INTERFACE_VERSION byte = 0x34

)
