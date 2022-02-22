package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([][]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		a[i] = make([]int, m)
	}
	// n行m個数を持つ2次元配列
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &a[i][j])
		}
	}
	// m行1個の配列
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &b[i])
	}

	for i := 0; i < n; i++ {
		c := 0
		for j := 0; j < m; j++ {
			c += a[i][j] * b[j]
		}
		fmt.Println(c)
	}
}
