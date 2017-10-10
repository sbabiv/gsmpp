package pdu

import (
	"sync/atomic"
	"time"
	"bytes"
	"encoding/binary"
)

const (
	HeaderLength = 16
	NullTerminated = "\x00"
)

var seq uint32 = uint32(time.Now().Unix())

type OptionalParameters map[uint16]*TLV

func coctet(v string) string {
	return v + NullTerminated
}

func sequenceInc() uint32 {
	return atomic.AddUint32(&seq,1)
}

type Header struct {
	Length uint32
	Id uint32
	Status uint32
	Sequence uint32
}

func NewHeader(len uint32, id uint32, status uint32, seq *uint32) *Header {
	var sequence uint32
	if seq != nil {
		sequence = *seq
	} else {
		sequence = sequenceInc()
	}

	return &Header{Length: len + HeaderLength, Id: id, Status: status, Sequence: sequence}
}

func (h *Header)Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, h)

	return buf.Bytes()
}

func (h *Header)GetBodyLen() uint32 {
	return h.Length - HeaderLength
}

type BindTransceiverResp struct {
	SystemId string
	OptionalParameters
}


type BindCommand struct {
	header *Header
	body   []byte
}

type UnbindCommand struct {
	*Header
}