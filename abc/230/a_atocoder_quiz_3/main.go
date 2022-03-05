// nolint
package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

func main() {
	var n int
	fmt.Scan(&n)

	ans := "AGC"
	if n >= 42 {
		ans += fmt.Sprintf("0%d", n+1)
	} else if utils.NumOfDigit(n) == 2 {
		ans += fmt.Sprintf("0%d", n)
	} else {
		ans += fmt.Sprintf("00%d", n)
	}
	fmt.Println(ans)
}
