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
	var name string
	var manager c.Manager

	fmt.Print("Enter user name: ")
	fmt.Fscan(os.Stdin, &name)
	// fmt.Println("Wait...")

	client := c.NewClient(network, address, &manager)
	settings := c.NewSettings(name)
	tui := c.NewTui(&manager)
	storage := c.NewStorage()

	manager = c.NewManager(client, settings, tui, storage)

	go client.Run()
	tui.Run()
}
