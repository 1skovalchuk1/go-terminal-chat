package client

import (
	"github.com/gorilla/websocket"
	"github.com/rivo/tview"
)

type Manager struct {
	tui      *Tui
	storage  *Storage
	client   *Client
	settings *Settings
}

type Settings struct {
	userName User
}

type Tui struct {
	tuiApp       *tview.Application
	board        *tview.TextView
	users        *tview.TextView
	inputMessage *tview.TextArea
	manager      *Manager
}

type Client struct {
	connection *websocket.Conn
	manager    *Manager
}

type Storage struct {
	messages []Message
	users    []User
}

type User string
type SentTime string
type ChatUsers string
type ChatMessages string

type Message struct {
	fromUser User
	time     SentTime
	text     string
}
