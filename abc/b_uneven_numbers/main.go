package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	ans := 0
	for i := 1; i <= N; i++ {
		// if numOfDigit(i)%2 != 0 {
		// 	ans++
		// }
		if i <= i && i <= 9 {
			ans++
		} else if 100 <= i && i <= 999 {
			ans++
		} else if 10000 <= i && i <= 99999 {
			ans++
		}
	}
	fmt.Println(ans)
}

func numOfDigit(n int) int {
	cnt := 0
	for n > 0 {
		n /= 10
		cnt++
	}
	return cnt
}
