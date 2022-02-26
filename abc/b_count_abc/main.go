package main

import (
	"fmt"
)

func main() {
	var N int
	var S string
	fmt.Scan(&N, &S)
	cnt := 0
	for i := 0; i < N-2; i++ {
		if S[i:i+2] == "ABC" {
			cnt++
		}
		// if strings.HasPrefix(S[i:], "ABC") {
		// 	cnt++
		// }
	}
	fmt.Println(cnt)
}
