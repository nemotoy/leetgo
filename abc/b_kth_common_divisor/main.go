package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

// AとBの公約数のうち、K番目に大きい整数
func main() {
	var A, B, K int
	fmt.Scan(&A, &B, &K)
	a := []int{}
	for i := 1; i <= utils.Min(A, B); i++ {
		if A%i == 0 && B%i == 0 {
			a = append(a, i)
		}
	}
	fmt.Println(a[K-1])
}
