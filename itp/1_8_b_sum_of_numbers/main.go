package main

import (
	"fmt"
	"nemotoy/leetgo/utils"
)

// For each dataset, print the sum of digits in x.
func main() {
	for {
		var x string
		fmt.Scan(&x)
		if x == "0" {
			break
		}
		fmt.Println(utils.SumAsStr(x))
	}
}

func main2() {
	for {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		fmt.Println(utils.Sum(x))
	}
}
