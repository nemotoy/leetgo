package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	- 3char毎の部分文字列
	- charsの各要素がユニークな総数。ただし、同じcharsはそれぞれカウントする。
*/
func countGoodSubstrings(s string) int {
	ret := 0
	for i := 0; i <= len(s)-3; i++ {
		c1, c2, c3 := s[i], s[i+1], s[i+2]
		if c1 == c2 || c2 == c3 || c3 == c1 {
			continue
		}
		ret++
	}
	return ret
}

func TestCountGoodSubstrings(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"xyzzaz", 1,
		},
		{
			"aababcabc", 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countGoodSubstrings(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
