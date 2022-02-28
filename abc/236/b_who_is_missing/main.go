package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	m := make(map[int]int)
	for i := 0; i < 4*n-1; i++ {
		var a int
		fmt.Scan(&a)
		m[a]++
	}
	for k, v := range m {
		if n+1 != v {
			fmt.Println(k)
		}
	}
}
