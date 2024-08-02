package server

import (
	"fmt"

	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Hub struct {
	clients       map[*Client]bool
	newClients    chan *Client
	logoutClients chan *Client
	messages      chan []byte
}

func (hub Hub) New() *Hub {
	hub.clients = make(map[*Client]bool)
	hub.newClients = make(chan *Client)
	hub.logoutClients = make(chan *Client)
	hub.messages = make(chan []byte)
	return &hub
}

func (hub *Hub) Run() {
	for {
		select {

		case message := <-hub.messages:
			hub.message(message)

		case client := <-hub.newClients:
			if hub.isValidUserName(client) {
				hub.newUserToCurrentUsers(client)
				hub.clients[client] = true
				hub.currentUsersToNewUser(client)

				fmt.Printf("send %v to current users\n", client.name)
				fmt.Printf("send current users to %v\n", client.name)
			} else {
				hub.existUserName(client)
				fmt.Printf("exist user name: %v\n", client.name)
			}

		case user := <-hub.logoutClients:
			user.closeChans()
			delete(hub.clients, user)
			hub.logoutUser(user)

			fmt.Printf("send logout user %v\n", user.name)

		}
	}
}

func (hub *Hub) message(message []byte) {
	for user := range hub.clients {
		select {
		case user.send <- message:
		default:
			user.closeChans()
			delete(hub.clients, user)
		}
	}
}
func (hub *Hub) isValidUserName(newClient *Client) bool {
	for user := range hub.clients {
		if user.name == newClient.name {
			return false
		}
	}
	return true
}

func (hub *Hub) existUserName(newClient *Client) {
	msg := message.Message{}.New(
		[]byte{},
		"",
		message.WarningExistClientName,
	)
	select {
	case newClient.send <- msg.ToBytes():
	default:
		newClient.closeChans()
		delete(hub.clients, newClient)
	}
}

func (hub *Hub) newUserToCurrentUsers(newClient *Client) {

	for otherUser := range hub.clients {
		msg := message.Message{}.New(
			[]byte(newClient.name),
			newClient.name,
			message.NewClient,
		)
		select {
		case otherUser.send <- msg.ToBytes():
		default:
			otherUser.closeChans()
			delete(hub.clients, otherUser)
		}
	}
}

func (hub *Hub) currentUsersToNewUser(newUser *Client) {
	currentUsersNames := ""
	for otherUser := range hub.clients {

		currentUsersNames += otherUser.name + "\n"
	}
	msg := message.Message{}.New(
		[]byte(currentUsersNames),
		newUser.name,
		message.UpdateClients,
	)
	select {
	case newUser.send <- msg.ToBytes():
	default:
		newUser.closeChans()
		delete(hub.clients, newUser)
	}
}

func (hub *Hub) logoutUser(logoutClient *Client) {

	for user := range hub.clients {
		msg := message.Message{}.New(
			[]byte(logoutClient.name),
			logoutClient.name,
			message.LogoutClient,
		)
		select {
		case user.send <- msg.ToBytes():
		default:
			user.closeChans()
			delete(hub.clients, user)
		}
	}
}
