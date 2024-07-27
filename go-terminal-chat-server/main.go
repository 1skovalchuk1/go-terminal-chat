package main

import (
	"fmt"
	"go-terminal-chat-server/src"
	"log"
	"net/http"
)

func main() {

	hub := src.Hub{}.New()
	go hub.Run()

	user := src.User{}.SetHub(hub)

	http.HandleFunc("/chat", user.Run)
	fmt.Println("Run server localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
