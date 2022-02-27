package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

func main() {
	var N int
	fmt.Scan(&N)
	// 初期値は約数の最大桁数になるNの桁数
	ans := utils.NumOfDigit(N)
	// N = A × B = B × A から A <= B を満たすAを列挙。
	// 1 <= A*A <= N から、 √1 <= A <= √N となり、時間計算量はO(√N)
	// ref. https://qiita.com/drken/items/872ebc3a2b5caaa4a0d0#%E4%BE%8B-10-%E7%B4%A0%E6%95%B0%E5%88%A4%E5%AE%9A-on
	for A := 1; A*A <= N; A++ {
		// N = A*B なので、AもBもNの約数
		if N%A != 0 {
			continue
		}
		var B = N / A
		cur := f(A, B)
		if ans > cur {
			ans = cur
		}
	}
	fmt.Println(ans)
}

// F(A, B) = AとBの桁数の大きい方
func f(a, b int) int {
	return utils.Max(utils.NumOfDigit(a), utils.NumOfDigit(b))
}
