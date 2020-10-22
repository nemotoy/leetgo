package main

import (
	"testing"
)

/*
	## summary
	部分文字列の最大長を返す。
*/
func maxPower(s string) int {
	var (
		res, tmp int
		prev     byte
	)
	for i := 0; i < len(s); i++ {
		if s[i] == prev {
			tmp += 1
		} else {
			prev = s[i]
			tmp = 1 // リセット値が1なのは、今回のindex分を加算するため
		}
		if res < tmp {
			res = tmp
		}
	}
	return res
}

func TestMaxPower(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"leetcode", 2,
		},
		{
			"abbcccddddeeeeedcba", 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := maxPower(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
