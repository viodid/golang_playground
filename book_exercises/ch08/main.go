package main

import (
	"fmt"
)

type IntOrFloat interface {
	~int | ~float64
}

func double[T IntOrFloat](n T) any {
	return n * 2
}

func main() {
	fmt.Println(double(2.3))
}
