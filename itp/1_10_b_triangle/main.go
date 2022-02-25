package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, C float64
	fmt.Scan(&a, &b, &C)

	S := a * b * math.Sin(math.Pi*C/180) / 2
	L := a + b + math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2)-2*a*b*math.Cos(math.Pi*C/180))
	h := b * math.Sin(math.Pi*C/180)
	fmt.Printf("%.8f\n", S)
	fmt.Printf("%.8f\n", L)
	fmt.Printf("%.8f\n", h)
}
