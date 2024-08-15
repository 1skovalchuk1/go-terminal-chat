package server

import "github.com/1skovalchuk1/go-terminal-chat/message"

type Hub struct {
	storage *Storage
	new     chan *Handler
	close   chan *Handler
	msg     chan [message.MessageSize]byte
}

func newHub() Hub {
	return Hub{
		storage: &Storage{handlers: make(map[*Handler]bool)},
		new:     make(chan *Handler, 100),
		close:   make(chan *Handler, 100),
		msg:     make(chan [message.MessageSize]byte, message.MessageSize*10),
	}
}

func (h *Hub) run() {
	for {
		select {
		case msg := <-h.msg:
			h.msgToAll(msg)
		case i := <-h.new:
			h.newUser(i)
		case i := <-h.close:
			h.closeHandler(i)
		}
	}
}

func (h *Hub) msgToAll(b [message.MessageSize]byte) {
	for i := range h.storage.handlers {
		select {
		case i.send <- b:
		default:
			h.closeHandler(i)
		}
	}
}

func (h *Hub) newUser(handler *Handler) {
	h.storage.addUser(handler, h.registry)
}

func (h *Hub) registry(handler *Handler) {
	users := handler.userName
	b := message.NewUser(handler.userName).ToBytes()
	for i := range h.storage.handlers {
		select {
		case i.send <- b:
			users += "\n" + i.userName

		default:
			h.closeHandler(i)
		}
	}
	msg := message.Users(users).ToBytes()
	select {
	case handler.send <- msg:
	default:
		h.closeHandler(handler)
	}
}

func (h *Hub) closeHandler(handler *Handler) {
	handler.close()
	delete(h.storage.handlers, handler)

	for i := range h.storage.handlers {
		msg := message.LogOut(handler.userName).ToBytes()
		select {
		case i.send <- msg:
		default:
			h.closeHandler(i)
		}
	}
}
