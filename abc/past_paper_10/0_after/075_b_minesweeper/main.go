// nolint
package main

import "fmt"

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	var s = make([][]rune, h)
	for i := 0; i < h; i++ {
		var str string
		fmt.Scan(&str)
		s[i] = []rune(str)
	}
	dx := [8]int{1, 0, -1, 0, 1, -1, -1, 1}
	dy := [8]int{0, 1, 0, -1, 1, 1, -1, -1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// 爆弾or空き
			if s[i][j] == '#' {
				continue
			}
			// 爆弾個数, 現在地は(i,j)
			cnt := 0
			for d := 0; d < 8; d++ {
				// 爆弾マス(ni,nj)
				ni, nj := i+dy[d], j+dx[d]
				// 高さの範囲外
				if ni < 0 || h <= ni {
					continue
				}
				// 幅の範囲外
				if nj < 0 || w <= nj {
					continue
				}
				if s[ni][nj] == '#' {
					cnt++
				}
			}
			s[i][j] = rune('0' + cnt)
		}
	}
	for i := 0; i < h; i++ {
		fmt.Println(string(s[i]))
	}
}
