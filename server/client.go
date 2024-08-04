package server

import (
	"fmt"
	"net"

	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Client struct {
	name       string
	connection net.Conn
	hub        *Hub
	send       chan []byte
}

func (client Client) Init(connection net.Conn, hub *Hub) Client {
	client.connection = connection
	client.hub = hub
	client.send = make(chan []byte)
	return client
}

func (client *Client) handleClientConnection() {
	go client.handleReadMessage()
	go client.handleWriteMessage()
}

func (client *Client) handleReadMessage() {
	defer client.connection.Close()

	for {
		dataBytes := make([]byte, 1024)
		_, err := client.connection.Read(dataBytes)
		if err != nil {
			client.hub.logoutClients <- client
			return
		}
		msg := message.Message{}.FromBytes(dataBytes)

		switch msg.TypeMsg {

		case message.TextMessage:
			client.hub.messages <- dataBytes
			continue

		case message.NewClient:
			client.setClientName(msg.FromClient)
			client.hub.newClients <- client
			continue

			// case message.LogoutClient:
			// 	client.hub.logoutClients <- client
			// 	continue
		}
	}
}

func (client *Client) handleWriteMessage() {
	defer client.connection.Close()

	for textMessage := range client.send {
		fmt.Println(client.name + ": ")
		fmt.Println("    " + string(message.Message{}.FromBytes(textMessage).DataBytes))
		client.sendMessage(textMessage)
	}
}

func (client *Client) setClientName(name string) {
	client.name = name
}

func (client *Client) sendMessage(message []byte) {
	_, err := client.connection.Write(message)
	if err != nil {
		fmt.Println("send message error: ", err)
		client.connection.Close()
	}
}

func (client *Client) closeChan() {
	fmt.Printf("close send chan: %v\n", client.name)
	close(client.send)
}
