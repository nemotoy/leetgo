// nolint
package main

import (
	"bufio"
	"fmt"
	"nemotoy/leetgo/utils"
	"os"
)

func main() {
	var str string
	var q int
	fmt.Scan(&str, &q)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < q; i++ {
		var order string
		var a, b int
		fmt.Scan(&order)
		fmt.Scan(&a, &b)
		switch order {
		case "print":
			fmt.Println(str[a : b+1])
		case "reverse":
			str = utils.ReverseSubString(str, a, b)
		case "replace":
			var p string
			fmt.Scan(&p)
			str = utils.ReplaceSubString(str, a, b, p)
		}
	}
}
