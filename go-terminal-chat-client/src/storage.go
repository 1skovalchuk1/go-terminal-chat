package src

func (s Storage) Init() *Storage {
	s.users = []User{}
	return &s
}

func (s *Storage) addMessage(message Message) {
	s.messages = append(s.messages, message)
}

func (s *Storage) addUser(newUser User) {
	s.users = append(s.users, newUser)
}

func (s *Storage) deleteUser(user User) {
	for i := range s.users {
		if s.users[i] == user {
			s.users = append(s.users[:i], s.users[i+1:]...)
			break
		}
	}
}
