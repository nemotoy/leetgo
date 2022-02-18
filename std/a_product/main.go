package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	aa, _ := strconv.Atoi(args[0])
	bb, _ := strconv.Atoi(args[1])
	if aa*bb%2 == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}
