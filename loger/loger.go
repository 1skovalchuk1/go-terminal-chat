package loger

import (
	"fmt"

	"github.com/1skovalchuk1/go-terminal-chat/message"
)

// ***Info***

func infoWrapper(msg string) message.Message {
	return message.Message{}.New(
		[]byte("***Info*** "+msg),
		"",
		message.Info,
	)
}

func InfoNewUser(userName string) message.Message {
	m := userName + " connected"
	return infoWrapper(m)
}

func InfoLogoutUser(userName string) message.Message {
	m := userName + " logout"
	return infoWrapper(m)
}
func InfoConnectUser() message.Message {
	m := "connected"
	return infoWrapper(m)
}

// *********
func InfoText(text string) message.Message {
	return infoWrapper(text)
}

// *********

// ***Error***

func errorWrapper(msg string) message.Message {
	return message.Message{}.New(
		[]byte(fmt.Sprint("***Error*** "+msg)),
		"",
		message.Info,
	)
}

func ErrorClientWrite() message.Message {
	return errorWrapper("client write")
}

func ErrorClientRead() message.Message {
	return errorWrapper("client read")
}

func ErrorClientConnection() message.Message {
	return errorWrapper("client connection")
}

func ErrorServerRead() message.Message {
	return errorWrapper("server read")
}

func ErrorServerWrite() message.Message {
	return errorWrapper("server write")
}
