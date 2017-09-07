package main

import (
	"fmt"
	"encoding/hex"
	"log"
	"github.com/gsmpp/smpp/pdu"
)



func main() {
	/*
	bind resp
		0000   c8 60 00 56 a3 3d 3c 94 d5 4f c9 81 08 00 45 00  .`.V.=<..O....E.
		0010   00 46 2f 80 40 00 37 06 35 bc 5f a3 6a 2f 0a 03  .F/.@.7.5._.j/..
		0020   0a a1 0a d8 c2 21 51 87 d2 2b 0e ca 85 91 50 18  .....!Q..+....P.
		0030   00 08 40 6e 00 00 00 00 00 1e 80 00 00 09 00 00  ..@n............
		0040   00 00 59 8c 7f 13 6d 6f 62 69 63 6f 6e 74 00 02  ..Y...mobicont..
		0050   10 00 01 34                                      ...4
	*/


	decoded, err := hex.DecodeString("0000001e8000000900000000598c7f136d6f6269636f6e74000210000134")
	if err != nil {
		log.Fatal(err)
	}

	pdu.Decode(decoded)


	fmt.Print("Implementation Protocol SMPP v3.4\n")
}
