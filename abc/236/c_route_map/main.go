package main

import "fmt"

// 急行が止まるか
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	// map[string]boolでもいい
	t := make([]string, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&t[i])
	}
	for _, sv := range s {
		res := "No"
		for _, tv := range t {
			if sv == tv {
				res = "Yes"
				break
			}
		}
		fmt.Println(res)
	}
}
