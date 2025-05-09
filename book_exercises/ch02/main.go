package main

import "fmt"

func main() {
	ex01()
	ex02()
	ex03()
}

func ex01() {
	i := 20
	var f = float32(i)
	fmt.Println(i, f)
}

func ex02() {
	const value = 42
	i := value
	var f float32 = value
	fmt.Println(i, f)
}

func ex03() {
	var b byte = 1<<8 - 1
	var smallI int32 = 1<<16 - 1
	var bigI uint64 = 1<<64 - 1
	fmt.Println(b, smallI, bigI)
}
