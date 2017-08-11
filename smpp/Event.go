package smpp

type ChannelEvents string

var (
	CONNECTED ChannelEvents = "CONNECTED"
	DISCONNECTED ChannelEvents = "DISCONNECTED"
	CONN_FAIL ChannelEvents = "CONNECTION FAILED"
	BIND ChannelEvents = "BIND"
	UNBIND ChannelEvents = "UNBIND"
	SEND_ENQUIRE_LINK ChannelEvents = "SEND_ENQUIRE_LINK"
)

type Event struct {
	ChannelEvents
	Message string
}

func NewEvent(events ChannelEvents, msg string) Event  {
	return Event{events, msg}
}