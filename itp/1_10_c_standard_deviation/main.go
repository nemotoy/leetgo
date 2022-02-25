package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var n int
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		var m, a float64
		s := make([]float64, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&s[i])
			m += s[i]
		}
		m /= float64(n)
		for i := 0; i < n; i++ {
			a += math.Pow(s[i]-m, 2)
		}
		a = math.Sqrt(a / float64(n))
		fmt.Printf("%.8f\n", a)
	}
}
