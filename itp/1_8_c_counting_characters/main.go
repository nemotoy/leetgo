package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// counts and reports the number of each alphabetical letter. Ignore the case of characters.
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
	text = strings.ToLower(text[:len(text)-1])

	var cnt [26]int
	for _, r := range text {
		if !unicode.IsSpace(r) {
			cnt[r-'a']++
		}
	}

	for i, c := range cnt {
		fmt.Printf("%c : %d\n", rune(i+'a'), c)
	}
}
