package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

func main() {
	var A, B, C, X, Y int
	fmt.Scan(&A, &B, &C, &X, &Y)
	// 仮の最大値
	ans := 100000 * 5000
	// 与件の個数の最大値
	l := 100000
	// i をABセットとすると 2Ci = max{x-k, 0} * A + max{y-k, 0} * B
	// 0 <= i <= l
	for i := 0; i <= l; i++ {
		ab := 2 * C * i
		a := utils.Max(X-i, 0) * A
		b := utils.Max(Y-i, 0) * B
		sum := ab + a + b
		if ans > sum {
			ans = sum
		}
	}
	fmt.Println(ans)
}
