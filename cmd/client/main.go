package main

import (
	"fmt"
	"net/url"
	"os"

	c "github.com/1skovalchuk1/go-terminal-chat/client"
)

const (
	scheme = "tcp"
	host   = "localhost:8080"
	path   = ""
)

func main() {
	url := url.URL{Scheme: scheme, Host: host, Path: path}

	var manager c.Manager
	var userName string

	client := c.Client{}.Init(&manager, url)
	storage := c.Storage{}.Init()
	tui := c.Tui{}.Init(&manager)
	settings := c.Settings{}.Init(userName)
	manager.Init(tui, client, storage, settings)

	for {
		fmt.Print("Enter user name: ")
		fmt.Fscan(os.Stdin, &userName)
		isRegister := client.Register(userName)
		if isRegister {
			settings.SetUserName(userName)
			break
		} else {
			fmt.Println("User name is already exist")
		}
	}

	client.Run()
	tui.Run()
}
