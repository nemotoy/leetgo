package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	a, b := s[0], s[2]
	// ai, bi := toInt(string(a)), toInt(string(b))
	ai, bi := int(a-'0'), int(b-'0')
	fmt.Println(ai * bi)
}

func toInt(s string) int {
	ai, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return ai
}
