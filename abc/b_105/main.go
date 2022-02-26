package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	ans := 0
	for i := 1; i <= N; i++ {
		if i%2 != 0 {
			divisor := numOfDivisor(i)
			if divisor == 8 {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func numOfDivisor(n int) int {
	divisor := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			divisor++
		}
	}
	return divisor
}
