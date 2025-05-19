package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	person1 := MakePerson("Bob", "Smith", 42)
	person2 := MakePerson("Alice", "Brown", 47)
	fmt.Println(person1, person2)
	ss := []string{"hey", "there"}
	s := "madame"
	UpdateSlice(ss, s)
	fmt.Println("after update", ss)
	GrowSlice(ss, s)
	fmt.Println("after grow", ss)
}

func MakePerson(fn, ln string, ag int) Person {
	return Person{fn, ln, ag}
}

func MakePersonPointer(fn, ln string, ag int) *Person {
	return &Person{fn, ln, ag}
}

func UpdateSlice(ss []string, s string) {
	ss[len(ss)-1] = s
	fmt.Println(ss)
}

func GrowSlice(ss []string, s string) {
	ss = append(ss, s)
	fmt.Println(ss)
}
