package main

import "fmt"

// Aji = Bij
// Aはh,wの二次元配列, Bはw,hの二次元配列
func main() {
	var h, w int
	fmt.Scan(&h, &w)
	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
		for j := 0; j < w; j++ {
			fmt.Scan(&a[i][j])
		}
	}
	b := make([][]int, w)
	for i := 0; i < w; i++ {
		b[i] = make([]int, h)
		for j := 0; j < h; j++ {
			b[i][j] = a[j][i]
		}
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			fmt.Printf("%d ", b[i][j])
			if j == h-1 {
				fmt.Println()
			}
		}
	}
}
