package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, x+1)
	}
	dp[0][0] = true
	var a, b int
	for i := 0; i < n; i++ {
		fmt.Scan(&a, &b)
		for j := 0; j <= x; j++ {
			if dp[i][j] {
				if j+a <= x {
					dp[i+1][j+a] = true
				}
				if j+b <= x {
					dp[i+1][j+b] = true
				}
			}
		}
	}
	if dp[n][x] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
