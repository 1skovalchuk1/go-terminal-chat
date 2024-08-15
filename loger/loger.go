package loger

import (
	"github.com/1skovalchuk1/go-terminal-chat/message"
)

// ***Info***

// *********
// func InfoText(text string) message.Message {
// 	return infoWrapper(text)
// }

// *********

// ***Error***

func errorWrapper(msg string) message.Message {
	return message.New(
		"***Error*** "+msg,
		"",
		message.InfoType,
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
