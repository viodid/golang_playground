package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(ex01())
}

func ex01() []int {
	out := make([]int, 0, 100)
	r := rand.New(rand.NewSource(42))
	for n := range 100 {
		rn := r.Int() % 100
		fmt.Println(n, rn)
		out = append(out, rn)
	}
	return out
}
