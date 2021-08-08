package main

import (
	"fmt"
	"testing"
)

/*
  ## summary
*/
func isAlienSorted(words []string, order string) bool {
	var orders [26]int
	for i, r := range order {
		orders[r-'a'] = i
	}

	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) {
				return false
			}
			if words[i][j] != words[i+1][j] {
				current := words[i][j] - 'a'
				next := words[i+1][j] - 'a'
				if orders[current] > orders[next] {
					return false
				}
				break
			}
		}
	}
	return true
}

func TestIsAlienSorted(t *testing.T) {
	tests := []struct {
		in  []string
		in2 string
		out bool
	}{
		{
			[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz", true,
		},
		{
			[]string{"word", "world", "row"}, "worldabcefghijkmnpqstuvxyz", false,
		},
		{
			[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz", false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %s", tt.in, tt.in2), func(t *testing.T) {
			got := isAlienSorted(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
