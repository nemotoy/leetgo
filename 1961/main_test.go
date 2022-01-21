package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	words[i]を連結した結果がsと同値か
*/
func isPrefixString(s string, words []string) bool {
	ele := ""
	for i := 0; i < len(words); i++ {
		ele += words[i]
		if s == ele {
			return true
		}
	}
	return false
}

func TestIsPrefixString(t *testing.T) {
	tests := []struct {
		in  string
		in2 []string
		out bool
	}{
		{
			"iloveleetcode",
			[]string{"i", "love", "leetcode", "apples"},
			true,
		},
		{
			"iloveleetcode",
			[]string{"apples", "i", "love", "leetcode"},
			false,
		},
		{
			"z",
			[]string{"z"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %v", tt.in, tt.in2), func(t *testing.T) {
			got := isPrefixString(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
