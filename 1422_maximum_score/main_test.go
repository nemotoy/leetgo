package main

import (
	"testing"
)

/*
	## summary
	左0,右1の数を加算した結果の最大値を返す。
*/
func maxScore(s string) int {
	res := 0
	for i := 1; i < len(s); i++ {
		// split s from index.
		left, right := s[:i], s[i:]
		// cnt is a counter.
		cnt := 0
		for _, r := range left {
			// when exists '0', increment cnt.
			if r == '0' {
				cnt++
			}
		}
		for _, r := range right {
			// when exists '1', increment cnt.
			if r == '1' {
				cnt++
			}
		}
		if res < cnt {
			res = cnt
		}
	}
	return res
}

func TestMaxScore(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"011101", 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := maxScore(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
