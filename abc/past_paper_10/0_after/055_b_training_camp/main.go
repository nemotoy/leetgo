package main

import (
	"fmt"
	"math"
	"nemotoy/leetgo/utils"
)

func main() {
	var n int
	fmt.Scan(&n)
	// i回目でi倍
	power := utils.Factorial(n)
	res := power % int(math.Pow10(9)+7) // 10の9乗
	fmt.Println(res)
}
