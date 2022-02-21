package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 任意行をスライスで受け取る
	scanner := bufio.NewScanner(os.Stdin)
	cnt := 1
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}
		fmt.Printf("case %d : %d\n", cnt, n)
		cnt++
	}
}
