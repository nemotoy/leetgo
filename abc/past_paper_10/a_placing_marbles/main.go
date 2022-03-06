package main

import (
	"fmt"
)

// 0 と 1 のみから成る 3 桁の番号 s に含まれる1の個数
func main() {
	var s string
	fmt.Scan(&s)
	var c int
	if s[0] == '1' {
		c++
	}
	if s[1] == '1' {
		c++
	}
	if s[2] == '1' {
		c++
	}
	fmt.Println(c)
}
