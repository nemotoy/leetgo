package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	res := "No"
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if i*j == N {
				res = "Yes"
				break
			}
		}
	}
	fmt.Println(res)
}
