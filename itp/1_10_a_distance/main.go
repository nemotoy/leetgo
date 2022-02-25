package main

import (
	"fmt"
	"math"
	"nemotoy/leetgo/utils"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)
	a, b := utils.Abs(x2-x1), utils.Abs(y2-y1)
	fmt.Printf("%.8f\n", math.Sqrt(float64(a*a+b*b)))
}
