package main

import (
	"fmt"
)

func main() {
	ex01()
	ex02()
}

func ex01() {
	greetings := []string{
		"Hello",
		"እው ሰላም ነው",
		"ဟယ်လို",
		"안녕하세요",
		"привіт",
		"Привет",
	}
	sub1 := greetings[:2]
	sub2 := greetings[2:4]
	sub3 := greetings[4:]
	fmt.Println(sub1, sub2, sub3)
}

func ex02() {
	message := "Hi 🧒 and 👩"
	mr := []rune(message)
	mb := []byte(message)
	fmt.Println(rune(message[3]), len(message))
	fmt.Printf("%v\n%v\n", mr, mb)
}
