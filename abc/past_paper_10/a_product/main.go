package main

import (
	"fmt"
)

// 2つの正整数a,bの積の偶奇判定
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	if (x*y)%2 == 0 {
		fmt.Println("Even")
		return
	}
	fmt.Println("Odd")
}
