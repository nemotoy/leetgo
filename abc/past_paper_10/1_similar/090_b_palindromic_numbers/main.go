package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

// a以上b以下の回文数の個数
func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := 0
	for i := a; i <= b; i++ {
		if utils.IsPalindromeInt(i) {
			ans++
		}
	}
	fmt.Println(ans)
}
