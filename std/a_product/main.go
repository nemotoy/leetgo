package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if (x*y)%2 == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}
