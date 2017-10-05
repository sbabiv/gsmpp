package text

import (
	"bytes"
	"unicode/utf16"
	"encoding/binary"
)

type Coding byte

const (
	Alphabet Coding = 0
	UCS2 			= 1 << 3
)

func ucs2(s string) []byte {
	var b bytes.Buffer
	for _, r := range utf16.Encode([]rune(s)) {
		u := make([]byte, 2, 2)
		binary.BigEndian.PutUint16(u, r)
		b.Write(u)
	}

	return b.Bytes()
}

func Encode(s string, c Coding) []byte {
	switch c {
	case Alphabet:
		return []byte(s)
	case UCS2:
		return ucs2(s)
	default:
		return []byte(s)
	}
}