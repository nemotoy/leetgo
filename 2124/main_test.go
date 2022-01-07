package main

import (
	"strings"
	"testing"
)

/*
	## summary
	"a"が"b"前に含まれる場合は真
*/
func checkString(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			after := s[i:]
			for i2 := 0; i2 < len(after); i2++ {
				if after[i2] == 'a' {
					return false
				}
			}
		}
	}
	return true
}

// "b"の次のindexに"a"がなければよい
func checkString2(s string) bool {
	return !strings.Contains(s, "ba")
}

func TestCheckString(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"aaabbb", true,
		},
		{
			"abab", false,
		},
		{
			"bbb", true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := checkString2(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
