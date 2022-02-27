package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	// a < b
	if b-a == 1 || b-a == 9 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
