# Doc

## Stdin

```go
// 1行固定
func main() {
	var x, y int
	fmt.Scan(&x, &y)
}

// 2行任意数
func main() {
	var n int
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
}

// 1とN個行固定数
func main() {
	var n int
	fmt.Scanf("%d", &n)
	t, x, y := []int{},[]int{},[]int{}
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i], &x[i], &y[i])
	}
}

// 1行を文字列で取得
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()
}
```
