package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	now, c := 0, 0
	for {
		if now == 1 {
			fmt.Println(c)
			break
		}
		if c >= n {
			fmt.Println(-1)
			break
		}
		c++
		now = a[now]
	}
}
