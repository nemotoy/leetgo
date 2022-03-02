package main

import (
	"fmt"
)

func main() {
	var s, t string
	fmt.Scan(&s, &t)

	// 先頭の差が他のindexの差と同値
	k := rune((t[0] - s[0] + 26) % 26)
	s2 := []rune(s)
	for i := 0; i < len(s); i++ {
		s2[i] = ((s2[i]-'a')+k)%26 + 'a'
	}

	if string(s2) == t {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func diffEle(a, b string, i int) rune {
	aa, bb := a[i], b[i]
	if aa < bb {
		return rune((bb - aa + 26) % 26)
	}
	return rune((aa - bb + 26) % 26)
}
