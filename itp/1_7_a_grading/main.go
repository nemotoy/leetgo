// nolint
package main

import (
	"bufio"
	"fmt"
	"os"
)

// reads a list of student test scores and evaluates the performance for each student.
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for {
		var m, f, r int
		fmt.Scan(&m, &f, &r)
		if m == -1 && f == -1 && r == -1 {
			break
		}
		sum := m + f
		if m == -1 || f == -1 {
			fmt.Println("F")
		} else if sum >= 80 {
			fmt.Println("A")
		} else if sum >= 65 && sum < 80 {
			fmt.Println("B")
		} else if sum >= 50 && sum < 65 {
			fmt.Println("C")
		} else if sum >= 30 && sum < 50 {
			if r >= 50 {
				fmt.Println("C")
			} else {
				fmt.Println("D")
			}
		} else {
			fmt.Println("F")
		}
	}
}
