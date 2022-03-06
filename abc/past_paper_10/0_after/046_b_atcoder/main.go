package main

import (
	"fmt"
)

// 隣り合うボールを別色に塗る場合の塗り方の個数
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	// 1つ目はk個
	ans := k
	for i := 1; i < n; i++ {
		// 左隣以外で塗るため、k-1個を最後尾まで実行する
		ans *= k - 1
	}
	fmt.Println(ans)
}
