package client

import "github.com/1skovalchuk1/go-terminal-chat/message"

type Storage struct {
	messages message.Messages
	users    Users
}

func (s Storage) Init() *Storage {
	s.users = Users{}
	s.messages = []message.Message{}
	return &s
}

func (s *Storage) addMessage(message message.Message) {
	s.messages = append(s.messages, message)
}

func (s *Storage) addUser(newUser User) {
	s.users = append(s.users, newUser)
}
func (s *Storage) addUsers(newUsers Users) {
	s.users = append(s.users, newUsers...)
}

func (s *Storage) deleteUser(user User) {
	for i := range s.users {
		if s.users[i] == user {
			s.users = append(s.users[:i], s.users[i+1:]...)
			break
		}
	}
}

func (s *Storage) deleteUsers() {
	s.users = []User{}
}
