package message

import "testing"

//49, 54, 58, 53, 54, 58, 51, 55

var msg = Message{
	TypeMsg: 5,
	Time:    [MessageTimeSize]byte{},
	From:    [MessageFromSize]byte{66, 105, 108, 108, 121},
	Body:    [MessageBodySize]byte{72, 101, 108, 108, 111, 32, 66, 105, 108, 108, 121},
}
var emptyMsg = Message{
	TypeMsg: 0,
	Time:    [MessageTimeSize]byte{},
	From:    [MessageFromSize]byte{},
	Body:    [MessageBodySize]byte{},
}

// Сheck a new message for equal
func TestNew_Equal(t *testing.T) {
	typeMsg := LogInType
	from := "Billy"
	body := "Hello Billy"

	want := msg
	res := New(body, from, typeMsg)

	if res != want {
		t.Errorf("Message.New_Equal is not equal\nres:= %v\nwant: %v\n", res, want)
	}
}

// Сheck the sequence and size of data in a message
func TestToBytes_Equal(t *testing.T) {
	want := [MessageSize]byte{msg.TypeMsg}
	s := want[:]
	s = append(s[:MessageTypeSize], msg.Time[:]...)
	s = append(s[:MessageTypeSize+MessageTimeSize], msg.From[:]...)
	_ = append(s[:MessageSize-MessageBodySize], msg.Body[:]...)

	res := msg.ToBytes()

	if res != want {
		t.Errorf("Message.TestToBytes_Equal is not equal\nres:= %v\nwant: %v\n", res, want)
	}
}

// Сheck parse for one message
func TestManyFromBytes_One_Equal(t *testing.T) {
	b := msg.ToBytes()

	res := ManyFromBytes(b[:])[0]
	want := msg

	if res != want {
		t.Errorf("Message.TestManyFromBytes_One_Equal is not equal\nres:= %v\nwant: %v\n", res, want)
	}
}

// Сheck parse for many messages
func TestManyFromBytes_Many_Equal(t *testing.T) {
	s := []byte{}
	b := msg.ToBytes()
	s = append(s, b[:]...)
	s = append(s, b[:]...)

	res := ManyFromBytes(s)
	want := []Message{msg, msg}

	if res[0] != want[0] || res[1] != want[1] {
		t.Errorf("Message.TestManyFromBytes_Many_Equal is not equal\nres:= %v\nwant: %v\n", res, want)
	}
}

// Сheck parse for first empty message
func TestManyFromBytes_First_Empty(t *testing.T) {

	b := emptyMsg.ToBytes()

	res := ManyFromBytes(b[:])

	if len(res) != 0 {
		t.Errorf("Message.TestManyFromBytes_One_Empty is not empty\nres:= %v\n", res)
	}
}

// Сheck messages length (when one message is empty, not first), and equality
func TestManyFromBytes_Other_Empty(t *testing.T) {
	s := []byte{}
	b1 := msg.ToBytes()
	b2 := emptyMsg.ToBytes()

	s = append(s, b1[:]...)
	s = append(s, b2[:]...)

	res := ManyFromBytes(s[:])
	resLen := len(ManyFromBytes(s[:]))

	wantLen := 1
	want := msg

	if len(res) != wantLen {
		t.Errorf("Message.TestManyFromBytes_One_Empty len is not equal\nlen(res):= %v\nwant: %v", resLen, wantLen)
	}
	if res[0] != msg {
		t.Errorf("Message.TestManyFromBytes_One_Empty message is not equal\nres:= %v\nwant: %v", res, want)
	}
}
