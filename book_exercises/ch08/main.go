package main

import (
	"fmt"
)

// == EX01 ==
type IntOrFloat interface {
	uint8 | uint16 | uint32 | uint64 |
		int8 | int16 | int32 | int64 |
		float32 | float64
}

func double[T IntOrFloat](n T) T {
	return n * 2
}

// == EX02 ==
type Printable interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
	fmt.Stringer
}

func printValue[P Printable](p P) {
	fmt.Println(p)
}

type myFloat float32

func (mf myFloat) String() string {
	return fmt.Sprintf("%f", mf)
}

func main() {
	var x float32 = 2.3
	var y int8 = 42
	fmt.Println(double(x), double(y))
	var z myFloat = 42.42
	printValue(z)
}
