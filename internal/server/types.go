package server

import (
	"github.com/gorilla/websocket"
)

const (
	textMessageType = byte(1)
	newUserType     = byte(2)
	updateUsersType = byte(3)
	logoutUserType  = byte(4)
	infoType        = byte(5)
)

type User struct {
	userName    string
	connection  *websocket.Conn
	hub         *Hub
	newUsers    chan []byte
	logoutUsers chan []byte
	updateUsers chan []byte
	messages    chan []byte
}

type Hub struct {
	users       map[*User]bool
	newUsers    chan *User
	logoutUsers chan *User
	messages    chan []byte
}
