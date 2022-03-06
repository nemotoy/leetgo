package main

import "fmt"

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	se, sw := [301010]int{}, [301010]int{}
	for i := 0; i < n; i++ {
		if s[i] == 'E' {
			se[i] = 1
		} else {
			sw[i] = 1
		}
	}

	for i := 1; i < n; i++ {
		sw[i] += sw[i-1]
		se[i] += se[i-1]
	}

	ans := 0
	for i := 0; i < n; i++ {
		sum := 0
		if i != 0 {
			sum += sw[i-1]
		}
		sum += se[n-1] - se[i]
		// min
		// TODO:
	}

	fmt.Println(ans)
}
