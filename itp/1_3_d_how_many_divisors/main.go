package main

import "fmt"

// prints the number of divisors of c between a and b
func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	cnt := 0
	for a <= b {
		if c%a == 0 {
			cnt++
		}
		a++
	}
	fmt.Println(cnt)
}
