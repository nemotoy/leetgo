// nolint
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	st, sh := 0, 0
	for i := 0; i < n; i++ {
		var t, h string
		fmt.Scan(&t, &h)
		if t > h {
			st += 3
		} else if t < h {
			sh += 3
		} else {
			st++
			sh++
		}
	}
	fmt.Println(st, sh)
}

func main2() {
	var n int
	fmt.Scan(&n)
	st, sh := 0, 0
	for i := 0; i < n; i++ {
		var t, h string
		fmt.Scan(&t, &h)
		if t == h {
			st++
			sh++
			continue
		}
		for i := 0; i < len(t); i++ {
			if t[i] > h[i] {
				st += 3
			} else if t[i] < h[i] {
				sh += 3
			} else {
				continue
			}
			break
		}
	}
	fmt.Println(st, sh)
}
