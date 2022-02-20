package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

// 区間の重なり
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	res := utils.Max(0, utils.Min(d, b)-utils.Max(c, a))
	fmt.Println(res)
}
