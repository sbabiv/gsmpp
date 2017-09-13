package tcp

import (
	"sync"
	"net"
	"errors"
)

type Client struct {
	sync.Mutex
	conn net.Conn
}

func (c *Client)Dial(network, address string) error {
	var err error
	c.conn, err = net.Dial(network, address)
	return err
}

func (c *Client) Write(b []byte) (n int, err error) {
	c.Lock()
	defer c.Unlock()
	return c.conn.Write(b)
}

func (c *Client) Read(len int) ([]byte, error){
	b := make([]byte, len, len)
	n, err := c.conn.Read(b)
	if err != nil || n < len{
		return nil, errors.New("Error read PDU")
	}

	return b, nil
}

func (c *Client) Close() {
	c.Lock()
	defer c.Unlock()
	c.conn.Close()
}

