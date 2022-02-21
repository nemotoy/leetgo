package main

import (
	"bufio"
	"fmt"
	"os"
)

// Draw a frame which has a height of H cm and a width of W cm.
// For example, the following figure shows a frame which has a height of 6 cm and a width of 10 cm.
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
				if i > 0 && i < h-1 && j > 0 && j < w-1 {
					res += "."
				} else {
					res += "#"
				}
			}
			res += "\n"
		}
		fmt.Println(res)
	}
}
