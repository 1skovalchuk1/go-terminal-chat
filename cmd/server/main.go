package main

import (
	"fmt"
	s "go-terminal-chat/internal/server"
	"log"
	"net/http"
)

func main() {

	hub := s.Hub{}.New()
	go hub.Run()

	user := s.User{}.SetHub(hub)

	http.HandleFunc("/chat", user.Run)
	fmt.Println("Run server localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
