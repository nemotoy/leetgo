package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	x, y, z := -1, -1, -1
	for xi := 0; xi <= a; xi++ {
		for yi := 0; yi <= a; yi++ {
			zi := a - xi - yi
			total := 10000*xi + 5000*yi + 1000*zi
			if (xi+yi+zi == a) && (total == b) {
				x, y, z = xi, yi, zi
				break
			}
		}
	}

	fmt.Println(x, y, z)
}
