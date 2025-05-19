package main

import "fmt"

const ITER int = 100_000_000

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println(CreateBigPersonSlice()[:1])
}
func CreateBigPersonSlice() []Person {
	s := make([]Person, 0, ITER)
	for _ = range ITER {
		s = append(s, Person{"Bob", "Smith", 42})
	}
	return s
}
