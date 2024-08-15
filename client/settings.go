package client

type Settings struct {
	name string
}

func NewSettings(name string) *Settings {
	return &Settings{
		name: name,
	}
}
