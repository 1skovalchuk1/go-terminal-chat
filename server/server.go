package server

import (
	"fmt"
	"net"
)

// Error handler
func e(err error) {
	if err != nil {
		panic(err)
	}
}

func Run() {
	laddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	e(err)
	listener, err := net.ListenTCP("tcp", laddr)
	fmt.Println("Run server: ", listener.Addr())
	e(err)

	hub := newHub()
	go hub.run()

	for {
		conn, err := listener.Accept()
		e(err)

		handler := newHandler(conn, &hub)

		handler.run()
	}
}
