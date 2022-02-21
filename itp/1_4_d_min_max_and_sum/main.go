package main

import (
	"fmt"
)

// reads a sequence of n integers ai(i=1,2,...n),
// and prints the minimum value, maximum value and sum of the sequence.
func main() {
	var n int
	fmt.Scanf("%d", &n)
	var min, max, sum int = 1000000, -1000000, 0
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		if tmp < min {
			min = tmp
		}
		if tmp > max {
			max = tmp
		}
		sum += tmp
	}
	fmt.Println(min, max, sum)
}
