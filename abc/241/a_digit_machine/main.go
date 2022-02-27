package main

import "fmt"

func main() {
	a := make([]int, 10)
	for i := 0; i < 10; i++ {
		fmt.Scan(&a[i])
	}
	ans := 0
	for i := 0; i < 3; i++ {
		view := a[ans]
		ans = view
	}
	fmt.Println(ans)
}
