// 中间人模式

package main

import "fmt"

// 中间人结构体
type ChatRoom struct{}

func (chatroom *ChatRoom) showMessage(user *User, message string) {
	fmt.Printf("%s say: %s\n", user.name, message)
}

// 实例需要交互的结构体
type User struct {
	name     string
	chatRoom *ChatRoom
}

func (user *User) say(message string) {
	user.chatRoom.showMessage(user, message)
}

func main() {
	chatRoom := &ChatRoom{}
	user1 := &User{
		name:     "Jack",
		chatRoom: chatRoom,
	}
	user2 := &User{
		name:     "Bob",
		chatRoom: chatRoom,
	}
	user3 := &User{
		name:     "Amy",
		chatRoom: chatRoom,
	}

	user1.say("Jack")
	user2.say("Bob")
	user3.say("Amy")
}
