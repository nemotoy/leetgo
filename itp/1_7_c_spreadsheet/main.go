package main

import (
	"fmt"
)

// reads the number of rows r, columns c and a table of r × c elements,
// and prints a new table, which includes the total sum for each row and column.
func main() {
	var r, c int
	fmt.Scan(&r, &c)
	tables := make([][]int, r+1)
	for i := 0; i < r+1; i++ {
		tables[i] = make([]int, c+1)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			// (i,j) table
			fmt.Scanf("%d", &tables[i][j])
			tables[r][j] += tables[i][j]
			tables[i][c] += tables[i][j]
			tables[r][c] += tables[i][j]
		}
	}
	for i := 0; i < r+1; i++ {
		for j := 0; j < c+1; j++ {
			if j == c {
				fmt.Printf("%d", tables[i][j])
			} else {
				fmt.Printf("%d ", tables[i][j])
			}
		}
		fmt.Println()
	}
}
