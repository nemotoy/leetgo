package main

import (
	"fmt"
	"strconv"
)

func main() {
	var abc string
	fmt.Scan(&abc)

	sa, sb, sc := abc[0], abc[1], abc[2]
	a := string(sa) + string(sb) + string(sc)
	b := string(sb) + string(sc) + string(sa)
	c := string(sc) + string(sa) + string(sb)

	ia, _ := strconv.Atoi(a)
	ib, _ := strconv.Atoi(b)
	ic, _ := strconv.Atoi(c)
	fmt.Println(ia + ib + ic)
}
