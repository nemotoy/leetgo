package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// print x and y in ascending order in a line
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		rx, ry := scanner.Text()[0], scanner.Text()[2]
		x, _ := strconv.Atoi(string(rx))
		y, _ := strconv.Atoi(string(ry))
		if x == 0 && y == 0 {
			break
		}
		if x > y {
			x, y = y, x
		}
		fmt.Printf("%d %d\n", x, y)
	}
}
