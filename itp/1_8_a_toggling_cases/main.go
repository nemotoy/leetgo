package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	// reads a first line as an string.
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	for _, r := range text {
		if unicode.IsUpper(r) {
			fmt.Printf("%s", strings.ToLower(string(r)))
		} else {
			fmt.Printf("%s", strings.ToUpper(string(r)))
		}
	}
	fmt.Println()
}
