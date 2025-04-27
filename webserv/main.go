package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("I hope I get the job!")
	f, err := os.Open("./message.txt")
	if err != nil {
		os.Exit(1)
	}
	var buffer byte
	f.Read(buffer)

}
