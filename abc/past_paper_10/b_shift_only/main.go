package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i]) // 1行スペース毎N個数
	}
	var c int
	for {
		var existOdd bool
		for _, v := range a {
			if v%2 != 0 {
				existOdd = true
				break
			}
		}
		if existOdd {
			break
		}

		for i := 0; i < n; i++ {
			a[i] /= 2
		}

		c++
	}
	fmt.Println(c)
}
