package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Draw a chessboard which has a height of H cm and a width of W cm.
// For example, the following figure shows a chessboard which has a height of 6 cm and a width of 10 cm.
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
			if i%2 == 0 {
				res += strings.Repeat("#.", w)
			} else {
				res += strings.Repeat(".#", w)
			}
			res += "\n"
		}
		fmt.Println(res)
	}
}
