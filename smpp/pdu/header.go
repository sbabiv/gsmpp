package pdu

import (
	"sync/atomic"
	"time"
	"bytes"
	"encoding/binary"
)

const HeaderLength = 4
const NullTerminated = "\x00"

var seq uint32 = uint32(time.Now().Unix())

func sequenceInc() uint32 {
	return atomic.AddUint32(&seq,1)
}

type Header struct {
	length uint32
	id uint32
	status uint32
	sequence uint32
}

func NewHeader(len uint32, id uint32, status uint32) *Header  {
	return &Header{length:len, id:id, status:status, sequence:sequenceInc()}
}

func (h *Header)Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, h)

	return buf.Bytes()
}

