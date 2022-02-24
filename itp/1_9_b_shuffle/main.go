package main

import (
	"fmt"
)

// reads a deck (a string) and a sequence of h, and prints the final state (a string).
func main() {
	for {
		var w string
		var m, h int
		fmt.Scan(&w)
		if w == "-" {
			break
		}
		fmt.Scan(&m)
		for i := 0; i < m; i++ {
			fmt.Scan(&h)
			w = w[h:] + w[:h]
		}
		fmt.Println(w)
	}
}
