package main

import "fmt"

func main() {
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	res := 0
	for ai := 0; ai <= a; ai++ {
		for bi := 0; bi <= b; bi++ {
			for ci := 0; ci <= c; ci++ {
				total := 500*ai + 100*bi + 50*ci
				if total == x {
					res++
				}
			}
		}
	}
	fmt.Println(res)
}
