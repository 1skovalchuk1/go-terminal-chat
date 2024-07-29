package server

import (
	"fmt"

	"github.com/1skovalchuk1/go-terminal-chat/message"

	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type User struct {
	userName   string
	connection *websocket.Conn
	hub        *Hub
	send       chan []byte
}

func (user User) Run(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade connection error: ", err)
		user.connection.Close()
	}
	user.send = make(chan []byte)
	user.connection = connection

	go user.ReadMessageHandler()
	go user.WriteMessageHandler()
}

func (user User) SetHub(hub *Hub) User {
	user.hub = hub
	return user

}

func (user *User) ReadMessageHandler() {
	defer func() {
		fmt.Printf("user %v logout\n", user.userName)
		user.hub.logoutUsers <- user
	}()
	for {
		_, dataBytes, err := user.connection.ReadMessage()
		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		msg := message.Message{}.FromBytes(dataBytes)

		switch msg.TypeMsg {

		case message.TextMessage:
			user.hub.messages <- dataBytes
			continue

		case message.NewUser:
			user.userName = msg.FromUser
			user.hub.newUsers <- user
			continue

		case message.LogoutUser:
			user.hub.logoutUsers <- user
			continue
		}
	}
}

func (user *User) WriteMessageHandler() {
	defer user.connection.Close()

	for textMessage := range user.send {
		user.sendMessage(textMessage)
		continue
	}
}

func (user *User) sendMessage(message []byte) {
	err := user.connection.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		fmt.Println("send message error: ", err)
		user.connection.Close()
	}
	fmt.Println("send message: ", string(message))
}

func (user *User) closeChans() {
	fmt.Printf("close and logout: %v\n", user.userName)
	close(user.send)
}
