package client

type User string
type Users []User

func (user User) fromBytes(b []byte) User {

	res := ""
	for _, i := range b {
		res += string(rune(i))
	}
	return User(res)
}

func (users Users) fromBytes(b []byte) []User {
	w := ""
	for _, i := range b {
		if rune(i) == '\n' {
			users = append(users, User(w))
			w = ""
			continue
		}
		w += string(rune(i))
	}
	return users
}

func (users Users) ChatUsers() string {
	res := ""
	for _, i := range users {
		res += string(i) + "\n"
	}
	return res
}
