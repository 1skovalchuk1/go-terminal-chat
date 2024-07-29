package main

import (
	"fmt"
	"net/url"
	"os"

	c "github.com/1skovalchuk1/go-terminal-chat/client"
)

const (
	scheme = "ws"
	host   = "localhost:8080"
	path   = "/chat"
)

func main() {

	url := url.URL{Scheme: scheme, Host: host, Path: path}

	var manager c.Manager

	var userName string

	fmt.Print("Enter user name: ")
	fmt.Fscan(os.Stdin, &userName)

	tui := c.Tui{}.Init(&manager)
	client := c.Client{}.Init(&manager, url)
	storage := c.Storage{}.Init()
	settings := c.Settings{}.Init(userName)
	manager = c.Manager{}.Init(tui, client, storage, settings)

	client.Run()
	tui.Run()
}
