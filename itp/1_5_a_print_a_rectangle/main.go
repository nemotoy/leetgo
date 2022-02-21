package main

import (
	"bufio"
	"fmt"
	"os"
)

// Draw a rectangle which has a height of H cm and a width of W cm.
// Draw a 1-cm square by single '#'.
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for {
		var h, w int
		fmt.Scan(&h, &w)
		if h == 0 && w == 0 {
			break
		}
		res := ""
		for i := 0; i < h; i++ {
			for j := 0; j < w; j++ {
				res += "#"
			}
			res += "\n"
		}
		fmt.Println(res)
	}
}
