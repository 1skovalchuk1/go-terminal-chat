package server

import (
	"fmt"
	"net"

	"github.com/1skovalchuk1/go-terminal-chat/message"
)

type Handler struct {
	userName string
	hub      *Hub
	conn     net.Conn
	send     chan [message.MessageSize]byte
}

func newHandler(conn net.Conn, hub *Hub) Handler {
	return Handler{
		conn: conn,
		send: make(chan [message.MessageSize]byte, message.MessageSize*10),
		hub:  hub,
	}
}

func (h *Handler) run() {
	go h.read()
	go h.write()
}

func (h *Handler) read() {
	for {
		// b := make([]byte, message.MessageSize)
		b := [message.MessageSize]byte{}
		_, err := h.conn.Read(b[:])
		e(err)
		msg := message.OneFromBytes(b)
		fmt.Println(
			"server read Type: ", fmt.Sprint(msg.TypeMsg), "\n",
			"server read Time: ", string(msg.Time[:]), "\n",
			"server read From: ", string(msg.From[:]), "\n",
			"server read Body: ", string((msg.Body)[:]))

		switch msg.TypeMsg {

		case message.TextType:
			h.hub.msg <- b

		case message.NewUserType:
			h.userName = msg.FromS()
			h.hub.new <- h

		case message.LogOutType:
			h.hub.close <- h
		}

	}
}

func (h *Handler) write() {
	for b := range h.send {
		msg := message.OneFromBytes(b)
		fmt.Println(
			"server write Type: ", fmt.Sprint(msg.TypeMsg), "\n",
			"server write Time: ", string(msg.TimeS()), "\n",
			"server write From: ", string(msg.FromS()), "\n",
			"server write Body: ", string(msg.BodyS()))
		_, err := h.conn.Write(b[:])
		e(err)
	}
}

func (h *Handler) close() {
	// TODO
}
