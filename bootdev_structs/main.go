package main

import (
	"fmt"
	"strconv"
)

type send interface {
	sendMsg() string
}

type messageToSend struct {
	phoneNumber int
	message     string
}

func (m messageToSend) sendMsg() string {
	return m.message + ": " + strconv.Itoa(m.phoneNumber)
}

func printMsg(s send) {
	fmt.Printf("%v\n", s.sendMsg())
}

func main() {
	printMsg(messageToSend{
		phoneNumber: 0,
		message:     "hey there!",
	})
}
