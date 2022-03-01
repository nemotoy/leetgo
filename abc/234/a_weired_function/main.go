package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	fmt.Println(f(f(f(t)+t) + f(f(t))))
}

func f(t int) int {
	return t*t + 2*t + 3
}
