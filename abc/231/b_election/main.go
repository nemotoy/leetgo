package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	ss := make(map[string]int)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		ss[s]++
	}

	max := 0
	ans := ""
	for k, v := range ss {
		if max < v {
			max = v
			ans = k
		}
	}
	fmt.Println(ans)
}
