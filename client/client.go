package client

import (
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Client struct {
	connection net.Conn
	manager    *Manager
}

func (client Client) Init(manager *Manager, url url.URL) *Client {

	connection, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("dial: ", err)
		os.Exit(1)
	}

	client.manager = manager
	client.connection = connection
	return &client
}

func (client *Client) sendMessage(byteMsg []byte) {
	_, err := client.connection.Write(byteMsg)
	println("Write")
	if err != nil {
		fmt.Println("client write error :", err)
		return
	}

}

func (client *Client) reciveMessage() {
	defer client.connection.Close()
	for {
		dataBytes := make([]byte, 1024)
		_, err := client.connection.Read(dataBytes)
		if err != nil {
			fmt.Println("client read err:", err)
			break
		}
		client.manager.reciveMessage(dataBytes)
	}
}

func (client *Client) Register(userName string) bool {
	msg := message.Message{}.New([]byte(userName), userName, message.NewClient)
	client.sendMessage(msg.ToBytes())
	dataBytes := make([]byte, 1024)
	_, err := client.connection.Read(dataBytes)
	if err != nil {
		fmt.Println("client read err:", err)
	}
	m := message.Message{}.FromBytes(dataBytes)

	if m.TypeMsg != message.WarningExistClientName {
		client.manager.reciveMessage(dataBytes)
	}

	return m.TypeMsg != message.WarningExistClientName
}

func (client *Client) Run() {

	go client.reciveMessage()
}

func (client *Client) close() {
	msg := message.Message{}.New([]byte{}, "", message.LogoutClient)
	client.connection.Write(msg.ToBytes())
	client.connection.Close()
}
