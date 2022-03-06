package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

// 区間の重なり。ない場合は0
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	// A: a-----|------|b
	// B:      c|------|------d
	// bとdの最小値から、aとcの最大値を引く
	res := utils.Max(0, utils.Min(d, b)-utils.Max(c, a))
	fmt.Println(res)
}
