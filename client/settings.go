package client

type Settings struct {
	userName string
}

func (s Settings) Init(userName string) *Settings {
	s.userName = userName
	return &s
}

func (s *Settings) SetUserName(userName string) {
	s.userName = userName
}
