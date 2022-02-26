package main

import (
	"fmt"
)

func main() {
	var S string
	fmt.Scan(&S)
	ans := 0
	// 前方・後方indexを取る
	for i := 0; i < len(S); i++ {
		for j := len(S) - 1; j > i; j-- {
			s := S[i:j]
			// 部分文字列が元文字列と同値は処理しない
			if s == S {
				continue
			}
			// 部分文字列の要件を検証
			var isSub bool = true
			for _, r := range s {
				if r != 'A' && r != 'C' && r != 'G' && r != 'T' {
					isSub = false
					break
				}
			}
			if isSub && len(s) > ans {
				ans = len(s)
			}
		}
	}
	fmt.Println(ans)
}
