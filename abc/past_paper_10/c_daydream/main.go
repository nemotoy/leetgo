package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)

	data := [4]string{"dream", "dreamer", "erase", "eraser"}
	var can bool
	for {
		var has bool = true
		for _, v := range data {
			if strings.HasSuffix(s, v) {
				s = s[len(v):]
				has = true
				can = true
				break
			} else {
				has = false
			}
		}
		if !has {
			break
		}
	}
	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
