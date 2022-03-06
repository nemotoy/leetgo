package main

import (
	"fmt"
)

// A op1 B op2 C op3 D = 7 となるように、op1,op2,op3 に + か - を入れた式を作る
func main() {
	var s string
	fmt.Scan(&s)

	a := int(s[0] - '0')
	b := int(s[1] - '0')
	c := int(s[2] - '0')
	d := int(s[3] - '0')

	if a+b+c+d == 7 {
		fmt.Printf("%d+%d+%d+%d=7", a, b, c, d)
	} else if a-b+c+d == 7 {
		fmt.Printf("%d-%d+%d+%d=7", a, b, c, d)
	} else if a-b-c+d == 7 {
		fmt.Printf("%d-%d-%d+%d=7", a, b, c, d)
	} else if a-b-c-d == 7 {
		fmt.Printf("%d-%d-%d-%d=7", a, b, c, d)
	} else if a+b-c+d == 7 {
		fmt.Printf("%d+%d-%d+%d=7", a, b, c, d)
	} else if a+b-c-d == 7 {
		fmt.Printf("%d+%d-%d-%d=7", a, b, c, d)
	} else if a-b+c-d == 7 {
		fmt.Printf("%d-%d-%d-%d=7", a, b, c, d)
	}
}
