package main

import (
	"fmt"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < q; i++ {
		var x, k int
		fmt.Scan(&x, &k)
		ans := -1
		for i, v := range a {
			if x == v {
				k--
			}
			if k == 0 {
				ans = i + 1
				break
			}
		}
		fmt.Println(ans)
	}
}
