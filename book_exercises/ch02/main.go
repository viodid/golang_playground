package main

import "fmt"

func main() {
	ex01()
	ex02()
}

func ex01() {
	i := 20
	var f float32 = float32(i)
	fmt.Println(i, f)
}

func ex02() {
	const value = 42
	i := value
	var f float32 = value
	fmt.Println(i, f)
}
