package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	for i := 0; i < q; i++ {
		var xq int
		fmt.Scan(&xq)
		num := sort.Search(n, func(i int) bool {
			return a[i] >= xq
		})
		fmt.Println(len(a) - num)
	}
}

func main2() {
	var n, q int
	fmt.Scan(&n, &q)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := q; i > 0; i-- {
		var xq int
		fmt.Scan(&xq)
		ans := 0
		for _, h := range a {
			if xq <= h {
				ans++
			}
		}
		fmt.Println(ans)
	}
}
