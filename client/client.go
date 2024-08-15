package client

import (
	"net"

	"github.com/1skovalchuk1/go-terminal-chat/e"
	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Client struct {
	conn    *net.TCPConn
	manager *Manager
	network string
	address string
}

func NewClient(network, address string, manager *Manager) *Client {
	return &Client{
		network: network,
		address: address,
		manager: manager,
	}
}

func (c *Client) receive() {
	for {
		b := make([]byte, message.MessageSize*10)
		_, err := c.conn.Read(b)
		if err != nil {
			c.conn.Close()
			c.Run()
			return
		}
		c.manager.receive(b)
	}
}

func (c *Client) send(byteMsg [message.MessageSize]byte) {
	_, err := c.conn.Write(byteMsg[:])
	e.Print(err)
}

func (c *Client) Run() {
	for {
		addr, err := net.ResolveTCPAddr(c.network, c.address)
		e.Print(err)
		conn, err := net.DialTCP(c.network, nil, addr)
		if err != nil {
			continue
		}
		c.conn = conn
		msg := message.NewUser(c.manager.name()).ToBytes()
		c.send(msg)
		c.receive()
		return
	}
}
