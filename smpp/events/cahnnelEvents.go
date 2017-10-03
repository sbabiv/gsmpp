package events

import (
	"time"
)

type ChannelEvents string

var (
	CONNECTED ChannelEvents = "CONNECTED"
	DISCONNECTED ChannelEvents = "DISCONNECTED"
	CONN_FAIL ChannelEvents = "CONNECTION FAILED"
	BOUND_TRX ChannelEvents = "BOUND_TRX"
	BOUND_RX ChannelEvents ="BOUND_RX"
	BOUND_TX ChannelEvents ="BOUND_TX"
	CLOSED ChannelEvents ="CLOSED"
	BIND_FAIL ChannelEvents = "BIND_FAILED"
	UNBIND ChannelEvents = "UNBIND"

	SEND_ENQUIRE_LINK ChannelEvents = "SEND_ENQUIRE_LINK"
	SEND_ENQUIRE_LINK_RESP ChannelEvents = "SEND_ENQUIRE_LINK_RESP"
	RECEIVE_ENQUITE_LINK ChannelEvents = "RECEIVE_ENQUITE_LINK"
	RECEIVE_ENQUITE_LINK_RESP ChannelEvents = "RECEIVE_ENQUITE_LINK_RESP"
	SEND_ENQUIRE_LINK_ERR ChannelEvents = "SEND_ENQUIRE_LINK_ERR"

	READ_PDU_ERR ChannelEvents = "READ_ERR"
	SKIP_PDU ChannelEvents = "SKIP_PDU"

	DELIVER_SM ChannelEvents = "DELIVER_SM"
)

type Event struct {
	ChannelEvents
	Time time.Time
}

func NewEvent(events ChannelEvents) Event  {
	return Event{events, time.Now()}
}