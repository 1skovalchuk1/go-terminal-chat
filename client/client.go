package client

import (
	"fmt"
	"net/url"
	"os"

	"github.com/1skovalchuk1/go-terminal-chat/message"
	"github.com/gorilla/websocket"
)

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

func (client Client) Init(manager *Manager, url url.URL) *Client {

	connection, _, err := websocket.DefaultDialer.Dial(url.String(), nil)

	if err != nil {
		fmt.Println("dial: ", err)
		os.Exit(1)
	}

	client.manager = manager
	client.connection = connection
	return &client
}

func (client *Client) sendMessage(byteMsg []byte) {
	err := client.connection.WriteMessage(websocket.TextMessage, byteMsg)
	if err != nil {
		fmt.Println("write:", err)
		return
	}

}

func (client *Client) reciveMessage() {
	defer client.connection.Close()
	for {
		_, message, err := client.connection.ReadMessage()
		if err != nil {
			fmt.Println("client read err:", err)
			break
		}
		client.manager.reciveMessage(message)
	}
}

func (client *Client) Run() {
	name := client.manager.settings.userName
	msg := message.Message{}.New([]byte(name), name, message.NewUser)
	client.sendMessage(msg.ToBytes())
	go client.reciveMessage()
}

func (client *Client) close() {
	client.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	client.connection.Close()
}
