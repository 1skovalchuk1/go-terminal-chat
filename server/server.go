package server

import (
	"fmt"
	"net"
)

type Server struct {
	network  string
	address  string
	hub      *Hub
	listener *net.Listener
}

func (server Server) Init(network string, address string, hub *Hub) Server {
	listener, err := net.Listen(network, address)
	fmt.Println("Run server: ", address)
	if err != nil {
		fmt.Println("Error TCP Server")
	}

	server.address = address
	server.network = network
	server.hub = hub
	server.listener = &listener
	return server
}

func (server Server) Run() {
	for {
		connection, err := (*server.listener).Accept()
		if err != nil {
			return
		}
		client := Client{}.Init(connection, server.hub)
		go client.handleClientConnection()
	}
}
