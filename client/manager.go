package client

import (
	"github.com/1skovalchuk1/go-terminal-chat/loger"
	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Manager struct {
	tui      *Tui
	storage  *Storage
	client   *Client
	settings *Settings
}

func (m *Manager) Init(tui *Tui, client *Client, storage *Storage, settings *Settings) *Manager {
	m.tui = tui
	m.client = client
	m.storage = storage
	m.settings = settings

	return m
}

func (m *Manager) reciveMessage(b []byte) {
	msg := message.Message{}.FromBytes(b)

	switch msg.TypeMsg {

	case message.TextMessage:
		m.addMessage(msg)
		return
	case message.UpdateClients:
		users := Users{}.fromBytes(msg.DataBytes)
		m.addUsers(users)
		//
		m.addMessage(loger.InfoNewUser("data: " + Users(users).ChatUsers()))
		m.addMessage(loger.InfoNewUser("storage: " + Users(m.storage.users).ChatUsers()))
		//
		return
	case message.LogoutClient:
		user := User("").fromBytes(msg.DataBytes)
		m.deleteUser(user)
		m.addMessage(loger.InfoLogoutUser(string(user)))
		m.updateAll()
		return
	case message.NewClient:
		user := User("").fromBytes(msg.DataBytes)
		m.addUser(user)
		m.addMessage(loger.InfoNewUser(string(user)))
		m.updateAll()
		return
	case message.Info:
		// m.storage.addMessage(msg)
		// chatMessages := m.storage.messages.ChatMessages()
		// m.tui.updateBoard(chatMessages)
		// return
	}
}

func (m *Manager) sendMessage(text string) {
	msg := message.Message{}.New([]byte(text), m.settings.userName, message.TextMessage)
	byteMsg := msg.ToBytes()
	m.client.send(byteMsg)
}

func (m *Manager) addMessage(msg message.Message) {
	m.storage.addMessage(msg)
	chatMessages := m.storage.messages.ChatMessages()
	m.tui.updateBoard(chatMessages)
}

func (m *Manager) addUser(user User) {
	m.storage.addUser(user)
	chatUsers := m.storage.users.ChatUsers()
	m.tui.updateUsers(chatUsers)
}

func (m *Manager) updateAll() {
	chatUsers := m.storage.users.ChatUsers()
	chatMessages := m.storage.messages.ChatMessages()
	m.tui.updateAll(chatUsers, chatMessages)
}

func (m *Manager) addUsers(users Users) {
	m.storage.addUsers(users)
	chatUsers := users.ChatUsers()
	m.tui.updateUsers(chatUsers)
}

func (m *Manager) deleteUser(user User) {
	m.storage.deleteUser(user)
}

func (m *Manager) deleteUsers() {
	m.storage.deleteUsers()
	chatUsers := m.storage.users.ChatUsers()
	m.tui.updateUsers(chatUsers)
}

func (m *Manager) close() {
	m.client.connection.Close()
	m.tui.close()
}
