package main

import (
	"fmt"
	"go-terminal-chat-client/src"
	"net/url"
	"os"
)

const (
	scheme = "ws"
	host   = "localhost:8080"
	path   = "/chat"
)

func main() {

	url := url.URL{Scheme: scheme, Host: host, Path: path}

	var manager src.Manager

	var userName src.User

	fmt.Print("Enter user name: ")
	fmt.Fscan(os.Stdin, &userName)

	tui := src.Tui{}.Init(&manager)
	client := src.Client{}.Init(&manager, url)
	storage := src.Storage{}.Init()
	settings := src.Settings{}.Init(userName)
	manager = src.Manager{}.Init(tui, client, storage, settings)

	client.Run()
	tui.Run()

}
