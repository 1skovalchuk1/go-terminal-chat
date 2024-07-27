package src

import (
	"fmt"
)

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
			closeUserChans(logoutUser)
			delete(hub.users, logoutUser)
			hub.logoutUser(logoutUser)

			fmt.Printf("send logout user %v\n", logoutUser.userName)

		}
	}
}

func (hub *Hub) message(message []byte) {
	for user := range hub.users {
		select {
		case user.messages <- message:
		default:
			closeUserChans(user)
			delete(hub.users, user)
		}
	}
}

func (hub *Hub) newUserToCurrentUsers(newUser *User) {

	for otherUser := range hub.users {
		select {
		case otherUser.newUsers <- []byte(newUser.userName):
		default:
			closeUserChans(otherUser)
			delete(hub.users, otherUser)
		}
	}
}

func (hub *Hub) currentUsersToNewUser(newUser *User) {
	currentUsersNames := ""
	for otherUser := range hub.users {
		currentUsersNames += otherUser.userName + "\n"
	}
	select {
	case newUser.updateUsers <- []byte(currentUsersNames):
	default:
		closeUserChans(newUser)
		delete(hub.users, newUser)
	}
}

func (hub *Hub) logoutUser(logoutUser *User) {

	for user := range hub.users {
		select {
		case user.logoutUsers <- []byte(logoutUser.userName):
		default:
			closeUserChans(user)
			delete(hub.users, user)
		}
	}
}
