package pdu

import (
	"sync/atomic"
	"time"
	"math/rand"
)

var seq uint32 = uint32(time.Now().Unix())

func init() {
	seq = rand.New(rand.NewSource(time.Now().UnixNano())).Uint32()
}


func sequenceInc() uint32 {
	return atomic.AddUint32(&seq,1)
}

type Header struct {
	length uint32
	id uint32
	status uint32
	sequence uint32
}

func NewHeader() *Header  {
	return &Header{length:16, id:1, status:1, sequence:sequenceInc()}
}

