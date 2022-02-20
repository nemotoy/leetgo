package main

import "fmt"

func main() {
	// 1と任意行
	var n int
	fmt.Scanf("%d", &n)
	m := map[int]int{}
	var d int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &d)
		m[d] = 0
	}
	fmt.Println(len(m))
}
