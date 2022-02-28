package main

import (
	"fmt"
	"math"
)

// -A <= N < B
func main() {
	var n float64
	fmt.Scan(&n)
	min, max := math.Pow(-2, 31), math.Pow(2, 31)
	if n >= min && n < max {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
