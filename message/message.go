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
	TextMessage = byte(1)
	NewUser     = byte(2)
	UpdateUsers = byte(3)
	LogoutUser  = byte(4)
	Info        = byte(5)
)

type Message struct {
	TypeMsg   byte
	FromUser  string
	Time      string
	DataBytes []byte
}

type Messages []Message

func (message *Message) Split() ([]byte, byte) {
	return message.DataBytes, message.TypeMsg
}

func (message *Message) setTime() {

	h, m, s := time.Now().Clock()
	receiptTime := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	message.Time = receiptTime
}

func (message Message) New(dataBytes []byte, fromUser string, typeMsg byte) Message {
	message.DataBytes = dataBytes
	message.TypeMsg = typeMsg
	message.FromUser = fromUser
	message.Time = ""
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
	message.setTime()
	return "[yellow]" + string(message.Time) +
		"  [red]" + string(message.FromUser) +
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
