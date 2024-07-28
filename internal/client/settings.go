package client

func (s Settings) Init(userName User) *Settings {
	// return Settings{
	s.userName = userName
	return &s
	// }
}
