package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, p string
	fmt.Scan(&s, &p)

	if isIn(s, p) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isIn(s, p string) bool {
	return strings.Contains(s+s, p)
}
