package src

import "fmt"

func splitMessage(message []byte) ([]byte, byte) {
	return message[1:], message[0]
}

func joinMessage(message []byte, typeMsg byte) []byte {
	return append([]byte{typeMsg}, message...)
}

func closeUserChans(user *User) {
	fmt.Printf("close and logout: %v\n", user.userName)
	close(user.newUsers)
	close(user.logoutUsers)
	close(user.updateUsers)
	close(user.messages)
}
