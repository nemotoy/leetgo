package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	m := map[int]int{}
	for i := 0; i < N; i++ {
		fmt.Scan(&i)
		m[i]++
	}
	fmt.Println(len(m))
}
