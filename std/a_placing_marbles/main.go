package main

import (
	"fmt"
)

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
