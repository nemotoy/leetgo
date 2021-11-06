package main

import (
	"fmt"
	"strings"
	"testing"
)

/*
  ## summary
*/
func numOfStrings(patterns []string, word string) int {
	ret := 0
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			ret++
		}
	}
	return ret
}

func TestNumOfStrings(t *testing.T) {
	tests := []struct {
		in  []string
		in2 string
		out int
	}{
		{[]string{"a", "abc", "bc", "d"}, "abc", 3},
		{[]string{"a", "b", "c"}, "aaaaabbbbb", 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %s", tt.in, tt.in2), func(t *testing.T) {
			got := numOfStrings(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
