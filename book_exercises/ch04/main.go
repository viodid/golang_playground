package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rs := ex01()
	fmt.Println(rs)
	ex02(rs)
	ex03()
}

func ex01() []int {
	out := make([]int, 0, 100)
	r := rand.New(rand.NewSource(42))
	for _ = range 100 {
		rn := r.Int() % 100
		out = append(out, rn)
	}
	return out
}

func ex02(rs []int) {
	for _, v := range rs {
		if v%2 == 0 && v%3 == 0 {
			fmt.Println("Six!", v)
		} else if v%2 == 0 {
			fmt.Println("Two!", v)
		} else if v%3 == 0 {
			fmt.Println("Three!", v)
		} else {
			fmt.Println("Never mind", v)
		}
	}
}

func ex03() {
	var total int
	for i := range 10 {
		total := i + total
		fmt.Println(total)
	}
	fmt.Println(total)

}
