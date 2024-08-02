package message

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

// Message types
const (
	TextMessage            = byte(1)
	NewClient              = byte(2)
	UpdateClients          = byte(3)
	LogoutClient           = byte(4)
	Info                   = byte(5)
	WarningExistClientName = byte(6)
)

type Message struct {
	TypeMsg    byte
	FromClient string
	Time       string
	DataBytes  []byte
}

type Messages []Message

func (message *Message) Split() ([]byte, byte) {
	return message.DataBytes, message.TypeMsg
}

func (message Message) New(dataBytes []byte, fromClient string, typeMsg byte) Message {
	h, m, s := time.Now().Clock()
	receiptTime := fmt.Sprintf("%02d:%02d:%02d", h, m, s)

	message.Time = receiptTime
	message.DataBytes = dataBytes
	message.TypeMsg = typeMsg
	message.FromClient = fromClient
	return message
}

func (message *Message) ToBytes() []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(message)
	if err != nil {
		log.Fatal("message encode error:", err)
	}
	return buf.Bytes()
}

func (message Message) FromBytes(b []byte) Message {
	buf := *bytes.NewBuffer(b)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(&message)
	if err != nil {
		log.Fatal("message decode error:", err)
	}
	return message
}

func (message *Message) chatMessage() string {

	if message.TypeMsg == Info {
		return "[yellow]" + string(message.Time) +
			"  [gray]" + string(message.DataBytes)
	}
	return "[yellow]" + string(message.Time) +
		"  [red]" + string(message.FromClient) +
		"[white]" +
		"\n          " + string(message.DataBytes)
}

func (messages Messages) ChatMessages() string {
	res := ""
	for _, msg := range messages {
		textMsg := msg.chatMessage()
		res += textMsg + "\n"
	}
	return res
}

// ***Info***

func infoMessage(message string) Message {
	return Message{}.New(
		[]byte(fmt.Sprint("***Info*** "+message)),
		"",
		Info,
	)
}

func InfoNewUser(userName string) Message {
	return infoMessage(fmt.Sprintf("%v connected", userName))

}
func InfoLogoutUser(userName string) Message {
	return infoMessage(fmt.Sprintf("%v logout", userName))
}
