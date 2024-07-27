package src

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func (user User) Run(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Print("upgrade connection error:", err)
		user.connection.Close()
	}
	user.newUsers = make(chan []byte)
	user.updateUsers = make(chan []byte)
	user.logoutUsers = make(chan []byte)
	user.messages = make(chan []byte)
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
		_, message, err := user.connection.ReadMessage()

		if err != nil {
			fmt.Println("read error:", err)
			break
		}

		byteMsg, requestType := splitMessage(message)

		switch requestType {

		case textMessageType:
			user.hub.messages <- byteMsg
			continue

		case newUserType:
			user.userName = string(byteMsg)
			user.hub.newUsers <- user
			continue

		case logoutUserType:
			user.hub.logoutUsers <- user
			continue
		}
	}
}

func (user *User) WriteMessageHandler() {
	defer func() {
		user.connection.Close()
	}()
	for {
		select {

		case textMessage, ok := <-user.messages:
			if !ok {
				user.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			user.sendMessage(textMessage, textMessageType)
			continue

		case newUserName, ok := <-user.newUsers:
			if !ok {
				user.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			user.sendMessage(newUserName, newUserType)
			continue

		case logoutUserName, ok := <-user.logoutUsers:
			if !ok {
				user.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			user.sendMessage(logoutUserName, logoutUserType)
			continue

		case currentUsersNames, ok := <-user.updateUsers:
			if !ok {
				user.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			user.sendMessage(currentUsersNames, updateUsersType)
			continue
		}
	}
}

func (user *User) sendMessage(byteMsg []byte, resposeType byte) {
	message := joinMessage(byteMsg, resposeType)
	err := user.connection.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		fmt.Println("send message error: ", err)
		user.connection.Close()
	}
	fmt.Println("send message: ", string(message))
}
