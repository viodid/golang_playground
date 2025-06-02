package main

import (
	"fmt"
)

type IntOrFloat interface {
	uint8 | uint16 | uint32 | uint64 |
		int8 | int16 | int32 | int64 |
		float32 | float64
}

func double[T IntOrFloat](n T) T {
	return n * 2
}

func main() {
	var x float32 = 2.3
	var y int8 = 42
	fmt.Println(double(x))
	fmt.Println(double(y))
}
