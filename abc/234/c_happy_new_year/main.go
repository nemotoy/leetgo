package main

import (
	"fmt"
	"strconv"
)

// 0と2からなる正の整数のうち、K番目に小さいもの
// 2進法のk番目を取得し、1を2に変換した結果を出力
func main() {
	var k int64
	fmt.Scan(&k)

	v := strconv.FormatInt(k, 2)
	ans := ""
	for i := 0; i < len(v); i++ {
		if v[i] == '1' {
			ans += "2"
		} else {
			ans += "0"
		}
	}
	fmt.Println(ans)
}
