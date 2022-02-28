package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	cur := 0
	for i := 0; i < n; i++ {
		var h int
		fmt.Scan(&h)
		if cur < h {
			cur = h
		} else {
			break
		}
	}
	fmt.Println(cur)
}
