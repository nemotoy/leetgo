package main

import "fmt"

// a 以上 b 以下の整数のうち、x で割り切れるものの個数
// f(n) が 0以上n以下の整数のうち、条件を満たす個数とすると、f(b)-f(a-1)
func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	fb := f(b, x)
	fa := f(a-1, x)
	ans := fb - fa
	fmt.Println(ans)
}

func f(n, x int) int {
	if n >= 0 {
		return n/x + 1
	}
	return 0
}
