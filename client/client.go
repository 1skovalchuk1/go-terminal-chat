package client

import (
	"net"
	"time"

	"github.com/1skovalchuk1/go-terminal-chat/loger"
	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Client struct {
	connection net.Conn
	manager    *Manager
	network    string
	address    string
}

func (client Client) Init(manager *Manager, network string, address string) *Client {

	connection, err := net.Dial(network, address)

	if err != nil {
		client.manager.addMessage(loger.ErrorClientConnection())
	}

	client.manager = manager
	client.connection = connection
	client.network = network
	client.address = address
	return &client
}

func (client *Client) send(byteMsg []byte) {
	_, err := client.connection.Write(byteMsg)
	if err != nil {
		client.manager.addMessage(loger.ErrorClientWrite())
	}
}

func (client *Client) reciveMessage() {
	for {
		time.Sleep(time.Second * 3) /// something wrong
		dataBytes := make([]byte, 1024)
		_, err := client.connection.Read(dataBytes)
		if err != nil {
			client.manager.addMessage(loger.ErrorClientRead())
			client.manager.deleteUsers()
			client.connection.Close()
			go client.autoconnect()
			break
		}
		client.manager.reciveMessage(dataBytes)
	}
}

func (client *Client) Register(userName string) bool {
	msg := message.Message{}.New([]byte(userName), userName, message.NewClient)
	client.send(msg.ToBytes())
	dataBytes := make([]byte, 1024)
	_, err := client.connection.Read(dataBytes)
	if err != nil {
		client.manager.addMessage(loger.ErrorClientRead())
		return false
	}
	m := message.Message{}.FromBytes(dataBytes)
	isNewClientName := m.TypeMsg != message.WarningExistClientName

	if isNewClientName {
		client.manager.reciveMessage(dataBytes)
	}

	return isNewClientName
}

func (client *Client) Run() {
	go client.reciveMessage()
}

func (client *Client) autoconnect() {
	for {
		time.Sleep(10 * time.Second)
		connection, err := net.Dial(client.network, client.address)
		if err != nil {
			client.manager.addMessage(loger.ErrorClientConnection())
			continue
		}
		client.manager.addMessage(loger.InfoConnectUser())
		client.connection = connection
		client.Register(client.manager.settings.userName)
		break
	}
	client.Run()

}
