package client

import "github.com/1skovalchuk1/go-terminal-chat/message"

type Manager struct {
	tui      *Tui
	client   *Client
	storage  *Storage
	settings *Settings
}

// create new Manager
func NewManager(client *Client, settings *Settings, tui *Tui, storage *Storage) Manager {
	return Manager{
		client:   client,
		settings: settings,
		tui:      tui,
		storage:  storage,
	}
}

// send message to client
func (m *Manager) send(text string) {
	msg := message.New(text, m.settings.name, message.TextType)
	b := msg.ToBytes()
	m.client.send(b)
}

// recieve message from client
func (m *Manager) receive(b []byte) {
	msgs := message.ManyFromBytes(b)
	for _, msg := range msgs {
		m.messageHandler(msg)
	}
}

// processes the message according to its type
func (m *Manager) messageHandler(msg message.Message) {
	switch msg.TypeMsg {
	case message.TextType:
		m.addMessage(msg)

	case message.NewUserType:
		m.addUser(msg)

	case message.LogOutType:
		m.removeUser(msg)

	case message.UpdateUsersType:
		m.addUsers(msg)

	case message.LogInType:
	case message.InfoType:
	}

}

// create text message
// func (m *Manager) message(text string) message.Message {
// 	return message.New(text, m.name(), message.TextType)
// }

// add message to chat
func (m *Manager) addMessage(msg message.Message) {
	msg.SetTime()
	m.storage.addMessage(msg)
	chatMessages := message.ToChatMessages(m.storage.messages)
	m.tui.updateBoard(chatMessages)
}

// add User to chat
func (m *Manager) addUser(msg message.Message) {
	user := msg.FromS()
	m.storage.addUser(user)
	chatUsers := message.ToChatUsers(m.storage.users)
	m.tui.updateUsers(chatUsers)
}

// add Users to chat
func (m *Manager) addUsers(msg message.Message) {
	users := msg.BodyS()
	m.storage.addUsers(users)
	chatUsers := message.ToChatUsers(m.storage.users)
	m.tui.updateUsers(chatUsers)
}

// remove User from chat
func (m *Manager) removeUser(msg message.Message) {
	user := msg.FromS()
	m.storage.removeUser(user)
	chatUsers := message.ToChatUsers(m.storage.users)
	m.tui.updateUsers(chatUsers)
}

// close chat
func (m *Manager) close() {
	m.client.conn.Close()
	m.tui.close()
}

// Getters

// get User name
func (m *Manager) name() string {
	return m.settings.name
}
