package main

import (
	"fmt"
	"os"

	c "github.com/1skovalchuk1/go-terminal-chat/client"
)

const (
	network = "tcp"
	address = "localhost:8080"
)

func main() {

	var manager c.Manager
	var userName string

	client := c.Client{}.Init(&manager, network, address)
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

	go client.Run()
	tui.Run()
}
