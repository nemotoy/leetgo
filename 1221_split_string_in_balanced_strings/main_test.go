package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	L = +1, R = -1とカウントし、0になった場合同数が出現した、すなわち分割単位と言える。
*/
func balancedStringSplit(s string) int {
	var res, c int
	for i := 0; i < len(s); i++ {
		if s[i] == 'L' {
			c += 1
		} else {
			c -= 1
		}
		if c == 0 {
			res += 1
		}
	}
	return res
}

func TestBalancedStringSplit(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"RLRRLLRLRL", 4, // RL, RRLL, RL, RL
		},
		{
			"RLLLLRRRLR", 3, // RL, LLLRRR, LR
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := balancedStringSplit(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
