package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	ans := 0
	for {
		if y <= x+10*ans {
			fmt.Println(ans)
			break
		}
		ans++
	}
}
