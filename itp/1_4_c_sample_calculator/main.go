package main

import (
	"bufio"
	"fmt"
	"os"
)

// reads two integers a, b and an operator op, and then prints the value of a op b.
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for {
		var a, b int
		var op string
		fmt.Scan(&a)
		fmt.Scan(&op)
		fmt.Scan(&b)
		var res int
		switch op {
		case "?":
			break
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			res = a / b
		}
		fmt.Println(res)
	}
}
