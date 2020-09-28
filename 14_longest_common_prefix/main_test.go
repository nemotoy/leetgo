package main

import (
	"fmt"
	"testing"
)

/*
	## summary
*/
func longestCommonPrefix(strs []string) string {
	return ""
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		in  []string
		out string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			[]string{"dog", "racecar", "car"},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := longestCommonPrefix(tt.in)
			if got != tt.out {
				t.Errorf("got: %s, want: %s", got, tt.out)
			}
		})
	}
}
