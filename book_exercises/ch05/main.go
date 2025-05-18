package main

import (
	"errors"
	"fmt"
	"strconv"
)

func main() {
	base3 := exponentBase(3)
	fmt.Println(base3(0))
	defer fmt.Println("defer executed")
	//log.Fatal("exit")
	fmt.Println(base3(4))
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, val := range expressions {
		n, err := calculator(val)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println(val, n)
	}
}

func exponentBase(base int) func(int) int {
	return func(exponent int) int {
		out := 1
		for _ = range exponent {
			out *= base
		}
		return out
	}
}

var opMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func calculator(val []string) (int, error) {

	if len(val) != 3 {
		return 0, errors.New("invalid expression")
	}
	op1, err := strconv.Atoi(val[0])
	if err != nil {
		return 0, err
	}
	op2, err := strconv.Atoi(val[2])
	if err != nil {
		return 0, err
	}
	f, ok := opMap[val[1]]
	if !ok {
		return 0, errors.New("invalid operand")
	}
	return f(op1, op2)
}

func add(x, y int) (int, error) { return x + y, nil }
func sub(x, y int) (int, error) { return x - y, nil }
func mul(x, y int) (int, error) { return x * y, nil }
func div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}
