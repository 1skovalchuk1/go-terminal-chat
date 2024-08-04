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
				hub.newClientToCurrentClients(client)
				hub.clients[client] = true
				hub.currentClientsToNewClient(client)
			} else {
				hub.existUserName(client)
				fmt.Printf("exist user name: %v\n", client.name)
			}

		case user := <-hub.logoutClients:
			user.closeChan()
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
			user.closeChan()
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
		newClient.closeChan()
		delete(hub.clients, newClient)
	}
}

func (hub *Hub) newClientToCurrentClients(newClient *Client) {

	currentClientsNames := ""
	for otherUser := range hub.clients {

		currentClientsNames += otherUser.name + "\n"
	}

	fmt.Println("newClientToCurrentClients:")
	fmt.Println("   currentClientsNames: ", currentClientsNames)
	fmt.Println("   newClient: ", newClient.name)

	for otherUser := range hub.clients {
		msg := message.Message{}.New(
			[]byte(newClient.name), // ???
			newClient.name,         // ???
			message.NewClient,
		)
		select {
		case otherUser.send <- msg.ToBytes():
		default:
			otherUser.closeChan()
			delete(hub.clients, otherUser)
		}
	}
}

func (hub *Hub) currentClientsToNewClient(newClient *Client) {
	currentClientsNames := ""
	for otherUser := range hub.clients {

		currentClientsNames += otherUser.name + "\n"
	}
	msg := message.Message{}.New(
		[]byte(currentClientsNames),
		newClient.name,
		message.UpdateClients,
	)
	fmt.Println("currentClientsToNewClient:")
	fmt.Println("   currentClientsNames: ", currentClientsNames)
	fmt.Println("   newClient: ", newClient.name)
	select {
	case newClient.send <- msg.ToBytes():
	default:
		newClient.closeChan()
		delete(hub.clients, newClient)
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
			user.closeChan()
			delete(hub.clients, user)
		}
	}
}
