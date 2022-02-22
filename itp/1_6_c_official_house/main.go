package main

import "fmt"

// Write a program which reads a sequence of tenant/leaver notices,
// and reports the number of tenants for each room.
func main() {
	var n int
	fmt.Scanf("%d", &n)
	var buildings [4][3][10]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			buildings[i][j] = [10]int{}
		}
	}
	for i := 0; i < n; i++ {
		var b, f, r, v int
		fmt.Scan(&b, &f, &r, &v)
		buildings[b-1][f-1][r-1] = v
	}
	for tenant, floors := range buildings {
		for _, rooms := range floors {
			for room, person := range rooms {
				fmt.Printf(" %d", person)
				if room == 9 {
					fmt.Println()
				}
			}
		}
		if tenant != 3 {
			fmt.Println("####################")
		}
	}
}
