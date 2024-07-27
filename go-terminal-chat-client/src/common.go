package src

import (
	"fmt"
	"time"
)

func toSentTime() SentTime {
	h, m, s := time.Now().Clock()
	res := fmt.Sprintf("%v:%v:%v", h, m, s)
	return SentTime(res)
}

func toMessage(text string, fromUser User) Message {
	return Message{
		fromUser: User(fromUser),
		time:     toSentTime(),
		text:     text,
	}
}

func messageToBytes(message Message) []byte {
	res := []byte(string(message.fromUser) + "\n" + string(message.time) + "\n" + string(message.text) + "\n")
	return res
}

func bytesToMessage(b []byte) Message {
	slice := []string{}
	w := ""

	for _, i := range b {
		if rune(i) == '\n' {
			slice = append(slice, w)
			w = ""
			continue
		}
		w += string(rune(i))
	}
	fromUser, time, text := slice[0], slice[1], slice[2]
	return Message{
		fromUser: User(fromUser),
		time:     SentTime(time),
		text:     text,
	}
}

func bytesToUser(b []byte) User {

	res := ""
	for _, i := range b {
		res += string(rune(i))
	}
	return User(res)
}

func bytesToUsers(b []byte) []User {
	res := []User{}
	w := ""
	// TODO easy to string
	for _, i := range b {
		if rune(i) == '\n' {
			res = append(res, User(w))
			w = ""
			continue
		}
		w += string(rune(i))
	}
	return res
}

func toChatUsers(users []User) ChatUsers {
	res := ""
	for _, i := range users {
		res += string(i) + "\n"
	}
	return ChatUsers(res)
}

func toChatMessages(messages []Message) ChatMessages {
	res := ""
	for _, msg := range messages {
		textMsg := "[yellow]" + string(msg.time) +
			"  [red]" + string(msg.fromUser) + "[white]" +
			"\n          " + string(msg.text)
		res += textMsg + "\n"
	}
	return ChatMessages(res)
}

func splitMessage(message []byte) ([]byte, byte) {
	return message[1:], message[0]
}

func joinMessage(message []byte, typeMsg byte) []byte {
	return append([]byte{typeMsg}, message...)
}
