package smpp

import (
	"sync"
	"net"
)

type tcpClient struct {
	sync.Mutex
	conn net.Conn
}

func (c *tcpClient)Dial(network, address string) error {
	var err error
	c.conn, err = net.Dial(network, address)

	return err
}

func (c *tcpClient) Write(b []byte) (n int, err error) {
	c.Lock()
	defer c.Unlock()

	return c.Write(b)
}

func (c *tcpClient) Read(b []byte) (n int, err error) {
	c.Lock()
	defer c.Unlock()

	return c.Write(b)
}

func (c *tcpClient) Close() {
	c.Lock()
	defer c.Unlock()
	c.conn.Close()
}

