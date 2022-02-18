package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	var c int
	for {
		var existOdd bool
		for i, v := range a {
			if v%2 != 0 {
				existOdd = true
			} else {
				a[i] = v / 2
			}
		}
		if existOdd {
			break
		}
		c++
	}
	fmt.Println(c)
}
