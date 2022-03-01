package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

func main() {
	var l, r int
	var s string
	fmt.Scan(&l, &r)
	fmt.Scan(&s)
	fmt.Println(utils.ReverseSubString(s, l-1, r-1))
}
