package message

import (
	"fmt"
	"time"
)

const (
	TextType                = byte(1)
	NewUserType             = byte(2)
	UpdateUsersType         = byte(3)
	LogOutType              = byte(4)
	LogInType               = byte(5)
	InfoType                = byte(6)
	WarningExistHandlerName = byte(7)
)

const (
	MessageTypeSize = 1
	MessageTimeSize = 8
	MessageFromSize = 128
	MessageBodySize = 1024
	MessageSize     = MessageTypeSize + MessageTimeSize + MessageFromSize + MessageBodySize
)

type Message struct {
	TypeMsg byte
	Time    [MessageTimeSize]byte
	From    [MessageFromSize]byte
	Body    [MessageBodySize]byte
}

func New(body string, from string, typeMsg byte) Message {

	bodyBytes := [MessageBodySize]byte{}
	fromBytes := [MessageFromSize]byte{}

	copy(bodyBytes[:], body)
	copy(fromBytes[:], from)

	// TODO check if string not too big

	return Message{
		TypeMsg: typeMsg,
		From:    fromBytes,
		Body:    bodyBytes,
	}
}

// parse  message to bytes
func (m Message) ToBytes() [MessageSize]byte {
	res := [MessageSize]byte{m.TypeMsg}
	s := res[:]
	s = append(s[:MessageTypeSize], m.Time[:]...)
	s = append(s[:MessageTypeSize+MessageTimeSize], m.From[:]...)
	_ = append(s[:MessageSize-MessageBodySize], m.Body[:]...)

	return res
}

// parse bytes to message
func OneFromBytes(b [MessageSize]byte) Message {
	m := Message{}

	m.TypeMsg = b[0]
	copy(m.Time[:], b[MessageTypeSize:MessageTypeSize+MessageTimeSize])
	copy(m.From[:], b[MessageTypeSize+MessageTimeSize:MessageSize-MessageBodySize])
	copy(m.Body[:], b[MessageSize-MessageBodySize:])

	return m
}

// parse bytes to slice of messages
func ManyFromBytes(bs []byte) []Message {
	byteMsg := [MessageSize]byte{}
	res := []Message{}
	if bs[0] == 0 {
		return res
	}
	for i, b := range bs {
		k := i % MessageSize
		if k == 0 && i != 0 {
			if b == 0 {
				break
			}
			byteMsg[k] = b
			msg := OneFromBytes(byteMsg)
			res = append(res, msg)
			byteMsg = [MessageSize]byte{}
		}
		byteMsg[k] = b
	}
	msg := OneFromBytes(byteMsg)
	res = append(res, msg)

	return res
}

// set time in message
func (m *Message) SetTime() {
	h, min, s := time.Now().Clock()
	t := fmt.Sprintf("%02d:%02d:%02d", h, min, s)
	copy(m.Time[:], t)
}

// *********** Messages ************

// create message new user
func NewUser(name string) Message {
	return New("", name, NewUserType)
}

// create message current users
func Users(users string) Message {
	return New(users, "", UpdateUsersType)
}

// create message logout user
func LogOut(name string) Message {
	return New("", name, LogOutType)
}

// ********* Info Messages *********

// create message info wrapper
func infoWrapper(msg string) Message {
	return New(
		"***Info*** "+msg,
		"",
		InfoType,
	)
}

// create message when new user add to chat
func InfoNewUser(userName string) Message {
	m := userName + " connected"
	return infoWrapper(m)
}

// create message when user logout from chat
func InfoLogoutUser(userName string) Message {
	m := userName + " disconnected"
	return infoWrapper(m)
}

// ************ Getters ************

// get message Time field as a string
func (m Message) TimeS() string {
	return byteToStr(m.Time[:])
}

// get message Body field as a string
func (m Message) BodyS() string {
	return byteToStr(m.Body[:])
}

// get message From field as a string
func (m Message) FromS() string {
	return byteToStr(m.From[:])
}

// ******** Parsers to chat ********

// parse one message to chat
func (m Message) ToChatMessage() string {

	if m.TypeMsg == InfoType {
		return "[yellow]" + m.TimeS() +
			"  [gray]" + m.BodyS()
	}
	return "[yellow]" + m.TimeS() +
		"  [red]" + m.FromS() +
		"[white]" +
		"\n          " + m.BodyS()
}

// parse messages to chat
func ToChatMessages(ms []Message) string {
	res := ""
	for _, i := range ms {
		textMsg := i.ToChatMessage()
		res += textMsg + "\n"
	}
	return res
}

// parse user to chat
func ToChatUser(user string) string {
	return "[red]" + user + "\n"
}

// parse users to chat
func ToChatUsers(users []string) string {
	res := ""
	for _, i := range users {
		res += ToChatUser(i)
	}
	return res
}

// *********************************

func byteToStr(b []byte) string {
	res := []byte{}
	for _, i := range b {
		if i == 0 {
			break
		}
		res = append(res, i)
	}
	return string(res)

}
