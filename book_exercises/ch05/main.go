package main

import (
	"fmt"
)

func exponentBase(base int) func(int) int {
	return func(exponent int) int {
		out := 1
		for _ = range exponent {
			out *= base
		}
		return out
	}
}

func main() {
	base3 := exponentBase(3)
	fmt.Println(base3(0))
	defer fmt.Println("defer executed")
	//log.Fatal("exit")
	fmt.Println(base3(4))
}
