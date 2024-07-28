package server

import (
	"github.com/gorilla/websocket"
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
