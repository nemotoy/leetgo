package main

import (
	"fmt"
	"testing"
)

/*
	## summary
*/
func maxDepth(s string) int {
	ret, cur := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cur++
			ret = max(ret, cur)
		} else if s[i] == ')' {
			cur--
		}
	}
	return ret
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"(1+(2*3)+((8)/4))+1", 3,
		},
		{
			"(1)+((2))+(((3)))", 3,
		},
		{
			"1+(2*3)/(2-1)", 1,
		},
		{
			"1", 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxDepth(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
