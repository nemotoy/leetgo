package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for {
		var n, x int
		fmt.Scan(&n, &x)
		if n == 0 && x == 0 {
			break
		}
		var cnt int
		for i := 0; i < n; i++ {
			for j := 1; j < n; j++ {
				for k := 2; j < n; j++ {
					if i+j+k == x {
						cnt++
					}
				}
			}
		}
		fmt.Println(cnt)
	}
}
