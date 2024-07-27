package src

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

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

func (client *Client) sendMessage(byteMsg []byte, requestType byte) {
	message := joinMessage(byteMsg, requestType)
	err := client.connection.WriteMessage(websocket.TextMessage, message)
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
	client.sendMessage([]byte(client.manager.settings.userName), newUserType)
	go client.reciveMessage()
}

func (client *Client) close() {
	client.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	client.connection.Close()
}
