package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	01の部分文字列の組み合わせ。
*/
func countBinarySubstrings(s string) int {
	cnt, pre, cur := 0, 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] != s[i-1] {
			cnt += min(pre, cur)
			pre, cur = cur, 1
		} else {
			cur++
		}
	}
	return cnt + min(pre, cur)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"00110011", 6,
		},
		{
			"10101", 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countBinarySubstrings(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
