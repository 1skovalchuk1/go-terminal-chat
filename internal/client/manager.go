package client

func (m Manager) Init(tui *Tui, client *Client, storage *Storage, settings *Settings) Manager {
	m.tui = tui
	m.client = client
	m.storage = storage
	m.settings = settings

	m.storage.addUser(m.settings.userName)
	users := toChatUsers(m.storage.users)
	m.tui.updateUsers(users)
	return m
}

// TODO error to board chat

func (m *Manager) sendMessage(text string) {
	msg := message(text, m.settings.userName)
	byteMsg := messageToBytes(msg)
	m.client.sendMessage(byteMsg, textMessageType)
}

func (m *Manager) reciveMessage(b []byte) {
	byteMsg, responseType := splitMessage(b)
	switch responseType {
	case textMessageType:
		msg := bytesToMessage(byteMsg)
		m.storage.addMessage(msg)
		messages := toChatMessages(m.storage.messages)
		m.tui.updateBoard(messages)
		return
	case updateUsersType:
		users := bytesToUsers(byteMsg)
		for _, u := range users {
			m.storage.addUser(u)
		}
		chatUsers := toChatUsers(m.storage.users)
		m.tui.updateUsers(chatUsers)
		return
	case logoutUserType:
		user := bytesToUser(byteMsg)
		m.storage.deleteUser(user)
		users := toChatUsers(m.storage.users)
		m.tui.updateUsers(users)
		return
	case newUserType:
		user := bytesToUser(byteMsg)
		m.storage.addUser(user)
		users := toChatUsers(m.storage.users)
		m.tui.updateUsers(users)
		return
	}
}

func (m *Manager) close() {
	m.client.close()
	m.tui.close()
}
