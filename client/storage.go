package client

import "github.com/1skovalchuk1/go-terminal-chat/message"

type Storage struct {
	messages []message.Message
	users    []string
}

func NewStorage() *Storage {
	return &Storage{
		messages: []message.Message{},
		users:    []string{},
	}
}

func (s *Storage) addMessage(message message.Message) {
	s.messages = append(s.messages, message)
}

func (s *Storage) addUser(newUser string) {
	s.users = append(s.users, newUser)
}

func (s *Storage) addUsers(newUsers string) {
	user := ""

	for _, i := range newUsers {
		if i == '\n' {
			s.addUser(user)
			user = ""
			continue
		}
		user += string(i)
	}
	s.addUser(user)
}

func (s *Storage) removeUser(user string) {
	for i := range s.users {
		if s.users[i] == user {
			s.users = append(s.users[:i], s.users[i+1:]...)
			break
		}
	}
}

// func (s *Storage) removeUsers() {
// 	s.users = []string{}
// }
