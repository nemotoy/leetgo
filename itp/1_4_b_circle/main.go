package main

import (
	"fmt"
	"math"
)

// Print the area and circumference of the circle in a line.
func main() {
	var r float64
	fmt.Scan(&r)
	p := math.Pi
	var area, circumference float64 = r * r * p, 2 * r * p
	fmt.Println(area, circumference)
}
