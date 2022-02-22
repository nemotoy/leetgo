package main

import (
	"fmt"
)

// reads a sequence and prints it in the reverse order.
func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Printf("%d", a[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
