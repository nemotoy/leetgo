package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanln(&n, &a, &b) // 1行スペース毎
	res := 0
	for i := 1; i <= n; i++ {
		sum := calcSumOfDigits(i)
		if sum >= a && sum <= b {
			res += i
		}
	}
	fmt.Println(res)
}

func calcSumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
