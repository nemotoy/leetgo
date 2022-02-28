package main

import (
	"fmt"
	"os"
)

// sの先頭に0個以上の"a"を付与して、回文になるか
func main() {
	// pref,sufから連続する'a'個数
	var x, y int
	var s string
	fmt.Scan(&s)
	l := len(s)

	for i := 0; i < l; i++ {
		if s[i] == 'a' {
			x++
		} else {
			break
		}
	}
	// すべて'a'の場合は真
	if x == l {
		fmt.Println("Yes")
		os.Exit(0)
	}
	for i := l - 1; i > 0; i-- {
		if s[i] == 'a' {
			y++
		} else {
			break
		}
	}
	// prefが長い場合、'a'を付与しても成立しない
	if x > y {
		fmt.Println("No")
		os.Exit(0)
	}
	// k = y-x
	for i := x; i < l-y; i++ {
		if s[i] != s[x+l-y-i-1] {
			fmt.Println("No")
			os.Exit(0)
		}
	}
	fmt.Println("Yes")
}
