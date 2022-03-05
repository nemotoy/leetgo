package main

import (
	"fmt"
	"strings"
)

// substring
func main2() {
	var s string
	fmt.Scan(&s)

	t := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < 5; j++ {
			t += "oxx"
		}
	}

	if strings.Contains(t, s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
