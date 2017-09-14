package tcp

import (
	"net"
)

type Client struct {
	conn net.Conn
}

func (c *Client)Dial(network, address string) error {
	var err error
	c.conn, err = net.Dial(network, address)
	return err
}

func (c *Client) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}

func (c *Client) Read(len int) ([]byte, error){
	b := make([]byte, len, len)
	_, err := c.conn.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

