package main

import "fmt"

// print the missing cards(52 cards include 13 ranks).
func main() {
	var n int
	fmt.Scan(&n)
	card := make([][]int, 4)
	mark := []string{"S", "H", "C", "D"}
	for i := 0; i < 4; i++ {
		card[i] = make([]int, 13)
	}
	for i := 0; i < n; i++ {
		var mk string
		var num int
		fmt.Scan(&mk, &num)
		for j := 0; j < 4; j++ {
			if mark[j] == mk {
				card[j][num-1] = 1
			}
		}
	}
	for j := 0; j < 4; j++ {
		for i := 0; i < 13; i++ {
			if card[j][i] == 0 {
				fmt.Printf("%s %d\n", mark[j], i+1)
			}
		}
	}
}
