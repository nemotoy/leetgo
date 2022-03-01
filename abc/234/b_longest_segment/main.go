package main

import (
	"fmt"
	"math"
)

// 2点間の最大値
func main() {
	var n int
	fmt.Scan(&n)

	xs := make([]int, n)
	ys := make([]int, n)
	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		xs[i] = x
		ys[i] = y
	}
	var ans float64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := xs[i] - xs[j]
			y := ys[i] - ys[j]
			ans = math.Max(ans, math.Sqrt(float64(x*x+y*y)))
		}
	}
	fmt.Printf("%.10f\n", ans)
}
