package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	res := "NO"
	for i := 0; i < b; i++ {
		if a*i%b == c {
			res = "YES"
		}
	}
	fmt.Println(res)
}
