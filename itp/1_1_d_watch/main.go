package main

import "fmt"

// 秒から時:分:秒への変換
func main() {
	var S int
	fmt.Scan(&S)
	h := S / 3600
	m := (S - (h * 3600)) / 60
	s := (S - m*60 - h*3600)
	fmt.Printf("%d:%d:%d", h, m, s)
}
