package client

import (
	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Manager struct {
	tui      *Tui
	storage  *Storage
	client   *Client
	settings *Settings
}

func (m *Manager) Init(tui *Tui, client *Client, storage *Storage, settings *Settings) Manager {
	m.tui = tui
	m.client = client
	m.storage = storage
	m.settings = settings

	return *m
}

// TODO error to board chat

func (m *Manager) sendMessage(text string) {
	msg := message.Message{}.New([]byte(text), m.settings.userName, message.TextMessage)
	byteMsg := msg.ToBytes()
	m.client.sendMessage(byteMsg)
}

func (m *Manager) reciveMessage(b []byte) {
	msg := message.Message{}.FromBytes(b)

	switch msg.TypeMsg {
	case message.TextMessage:
		m.storage.addMessage(msg)
		chatMessages := m.storage.messages.ChatMessages()
		m.tui.updateBoard(chatMessages)
		return
	case message.UpdateClients:
		users := Users{}.fromBytes(msg.DataBytes)
		for _, user := range users {
			m.storage.addUser(User(user))
		}
		chatUsers := m.storage.users.ChatUsers()
		m.tui.updateUsers(chatUsers)
		return
	case message.LogoutClient:
		user := User("").fromBytes(msg.DataBytes)
		m.storage.deleteUser(user)
		m.storage.addMessage(message.InfoLogoutUser(string(user)))
		chatUsers := m.storage.users.ChatUsers()
		chatMessages := m.storage.messages.ChatMessages()
		m.tui.updateAll(chatUsers, chatMessages)
		return
	case message.NewClient:
		user := User("").fromBytes(msg.DataBytes)
		m.storage.addUser(user)
		m.storage.addMessage(message.InfoNewUser(string(user)))
		chatUsers := m.storage.users.ChatUsers()
		chatMessages := m.storage.messages.ChatMessages()
		m.tui.updateAll(chatUsers, chatMessages)
		return
	case message.Info:
		// m.storage.addMessage(msg)
		// chatMessages := m.storage.messages.ChatMessages()
		// m.tui.updateBoard(chatMessages)
		// return
	}
}

func (m *Manager) close() {
	m.client.close()
	m.tui.close()
}
