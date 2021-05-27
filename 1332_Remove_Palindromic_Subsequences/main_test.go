package main

import (
	"fmt"
	"testing"
)

func removePalindromeSub(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	i, j := 0, l-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return 2
		}
	}
	return 1
}

func TestRemovePalindromeSub(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"ababa", 1,
		},
		{
			"abb", 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := removePalindromeSub(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
