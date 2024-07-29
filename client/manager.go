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

func (m Manager) Init(tui *Tui, client *Client, storage *Storage, settings *Settings) Manager {
	m.tui = tui
	m.client = client
	m.storage = storage
	m.settings = settings

	m.storage.addUser(User(m.settings.userName))
	chatUsers := m.storage.users.ChatUsers()
	m.tui.updateUsers(chatUsers)
	return m
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
	case message.UpdateUsers:
		users := Users{}.fromBytes(msg.DataBytes)
		for _, user := range users {
			m.storage.addUser(User(user))
		}
		chatUsers := m.storage.users.ChatUsers()
		m.tui.updateUsers(chatUsers)
		return
	case message.LogoutUser:
		user := User("").fromBytes(msg.DataBytes)
		m.storage.deleteUser(user)
		chatUsers := m.storage.users.ChatUsers()
		m.tui.updateUsers(chatUsers)
		return
	case message.NewUser:
		user := User("").fromBytes(msg.DataBytes)
		m.storage.addUser(user)
		chatUsers := m.storage.users.ChatUsers()
		m.tui.updateUsers(chatUsers)
		return
	}
}

func (m *Manager) close() {
	m.client.close()
	m.tui.close()
}
