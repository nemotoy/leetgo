package main

import "fmt"

func main() {
	var s string
	var a, b int
	fmt.Scan(&s)
	fmt.Scan(&a, &b)
	as, bs := s[a-1], s[b-1]
	out := []rune(s)
	out[a-1] = rune(bs)
	out[b-1] = rune(as)
	fmt.Println(string(out))
}
