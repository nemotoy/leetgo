package main

import (
	"fmt"
	"strings"
)

// reads a word W and a text T, and prints the number of word W which appears in text T
func main() {
	var w string
	fmt.Scan(&w)
	cnt := 0
	for {
		var text string
		fmt.Scan(&text)
		if text == "END_OF_TEXT" {
			break
		}
		for _, s := range strings.Split(text, " ") {
			if w == strings.ToLower(s) {
				cnt++
			}
		}
	}
	fmt.Println(cnt)
}
