package main

import (
	"fmt"
	"net"
	"encoding/hex"
)

func main(){
	//0000002F
	/*const s = "00000002"
	b, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	var v = 0x00000009
	fmt.Println(v)

	result := binary.BigEndian.Uint32(b)
	fmt.Println(result)
	fmt.Printf("%s\n", b)

	//0000002F000000020000000000000001534D50503354455354007365637265743038005355424D4954310000010100 hex => byte[] len 47
	b, _ = hex.DecodeString("0000002F000000020000000000000001534D50503354455354007365637265743038005355424D4954310000010100")
	l := len(b)
	fmt.Printf("%v", l)

	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 2)


	str := hex.EncodeToString(bs)
	fmt.Println("hex:")
	fmt.Println(str)


	fmt.Println(bs)*/

	/*b, _ := hex.DecodeString("736563726574303800")
	s := string(b[:])
	fmt.Println(s)

	src := []byte("secret08")
	src = append(src, 0)
	val := hex.EncodeToString(src)
	fmt.Println(val)*/

	//listen()

	//client()

	fmt.Print("Start Implementation Protocol SMPP v3.4")

}


func client(){
	/*conn, err := net.Dial("tcp", "lk.rapporto.ru:2776")*/
	conn, err := net.Dial("tcp", "localhost:8075")
	if err != nil {
		fmt.Println("conn error ")
	}

	for {
		b, _ := hex.DecodeString("0000002F000000020000000000000001534D50503354455354007365637265743038005355424D4954310000010100")
		n, err := conn.Write(b)
		if err != nil {
			fmt.Println("write err: " + err.Error())
		}

		resp := make([]byte, 16)
		n, err = conn.Read(resp)
		if err != nil {
			fmt.Println("read error: " +  err.Error())
		}

		fmt.Println(n)
	}

	//conn.Close()

	//status, err := bufio.NewReader(conn).
}

func listen(){
	ln, err := net.Listen("tcp", ":8075")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("client accept")
		for {
			req := make([]byte, 1024)
			conn.Read(req)

			fmt.Println(string(req))
			conn.Write([]byte("success 200"))
		}
	}
}

//bind_transceiver 0x00000009

/*
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, 2)
	fmt.Println(bs)
*/

/*
 smsc
 	Хост: lk.rapporto.ru
	Порт: 2776

    Логин (system_id): wildberries2
	Пароль	Wild24BS

*/