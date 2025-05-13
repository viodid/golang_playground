package main

import (
	"fmt"
)

func main() {
	ex01()
	ex02()
	ex03()
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

type Employee struct {
	firstName string
	lastName  string
	id        uint
}

func ex03() {
	i1 := Employee{
		"Ana",
		"Bohuelo",
		1,
	}
	i2 := Employee{
		firstName: "Ana",
		lastName:  "Bohuelo",
		id:        1,
	}
	var i3 Employee
	i3.firstName = "Ana"
	i3.lastName = "Bohuelo"
	i3.id = 1
	fmt.Println(i1, i2, i3)
}
