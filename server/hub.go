package server

import (
	"fmt"

	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Hub struct {
	users       map[*User]bool
	newUsers    chan *User
	logoutUsers chan *User
	messages    chan []byte
}

func (hub Hub) New() *Hub {
	hub.users = make(map[*User]bool)
	hub.newUsers = make(chan *User)
	hub.logoutUsers = make(chan *User)
	hub.messages = make(chan []byte)
	return &hub
}

func (hub *Hub) Run() {
	for {
		select {

		case message := <-hub.messages:
			hub.message(message)

			fmt.Printf("send '%v' to users\n", string(message))

		case newUser := <-hub.newUsers:
			hub.newUserToCurrentUsers(newUser)
			hub.currentUsersToNewUser(newUser)
			hub.users[newUser] = true

			fmt.Printf("send %v to current users\n", newUser.userName)
			fmt.Printf("send current users to %v\n", newUser.userName)

		case logoutUser := <-hub.logoutUsers:
			logoutUser.closeChans()
			delete(hub.users, logoutUser)
			hub.logoutUser(logoutUser)

			fmt.Printf("send logout user %v\n", logoutUser.userName)

		}
	}
}

func (hub *Hub) message(message []byte) {
	for user := range hub.users {
		select {
		case user.send <- message:
		default:
			user.closeChans()
			delete(hub.users, user)
		}
	}
}

func (hub *Hub) newUserToCurrentUsers(newUser *User) {

	for otherUser := range hub.users {
		msg := message.Message{}.New(
			[]byte(newUser.userName),
			newUser.userName,
			message.NewUser,
		)
		select {
		case otherUser.send <- msg.ToBytes():
		default:
			otherUser.closeChans()
			delete(hub.users, otherUser)
		}
	}
}

func (hub *Hub) currentUsersToNewUser(newUser *User) {
	currentUsersNames := ""
	for otherUser := range hub.users {
		currentUsersNames += otherUser.userName + "\n"
	}
	msg := message.Message{}.New(
		[]byte(currentUsersNames),
		newUser.userName,
		message.UpdateUsers,
	)
	select {
	case newUser.send <- msg.ToBytes():
	default:
		newUser.closeChans()
		delete(hub.users, newUser)
	}
}

func (hub *Hub) logoutUser(logoutUser *User) {

	for user := range hub.users {
		msg := message.Message{}.New(
			[]byte(logoutUser.userName),
			logoutUser.userName,
			message.LogoutUser,
		)
		select {
		case user.send <- msg.ToBytes():
		default:
			user.closeChans()
			delete(hub.users, user)
		}
	}
}
