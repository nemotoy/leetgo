package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	var x, y int
	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			x += a[i]
		} else {
			y += a[i]
		}
	}
	fmt.Println(x - y)
}
