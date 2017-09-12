package pdu

type status int

const (
	//no error
	ESME_ROK status = 0x00000000

	//Message Length is invalid
	ESME_RINVMSGLEN status = 0x00000001

	//Command Length is invalid
	ESME_RINVCMDLEN status = 0x00000002

	//Invalid Command ID
	ESME_RINVCMDID status = 0x00000003

	//Incorrect BIND Status for given command
	ESME_RINVBNDSTS status = 0x00000004

	//ESME Already in Bound State
	ESME_RALYBND status = 0x00000005

	//Invalid Priority Flad
	ESME_RINVPRTFLG status = 0x00000006

	//Invalid Registred Delivery Flag
	ESME_RINVREGDLVFLG = 0x00000007

	//System Error
	ESME_RSYSERR status = 0x00000008
)


