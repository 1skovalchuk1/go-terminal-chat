package client

func (s Settings) Init(userName User) *Settings {
	s.userName = userName
	return &s
}
