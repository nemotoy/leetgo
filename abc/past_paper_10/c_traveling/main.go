package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	t, x, y := [110000]int{}, [110000]int{}, [110000]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i+1], &x[i+1], &y[i+1])
	}
	var can bool = true

	for i := 0; i < n; i++ {
		dt := t[i+1] - t[i]
		dist := abs(x[i+1]-x[i]) + abs(y[i+1]-y[i])
		if dt < dist {
			can = false
		}
		if dist%2 != dt%2 {
			can = false
		}
	}

	if can {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func abs(x int) int {
	if x < 1 {
		return -x
	}
	return x
}
