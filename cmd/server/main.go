package main

import (
	s "github.com/1skovalchuk1/go-terminal-chat/server"
)

const (
	network = "tcp"
	address = "localhost:8080"
)

func main() {

	hub := s.Hub{}.New()
	go hub.Run()

	server := s.Server{}.Init(network, address, hub)
	server.Run()
}
