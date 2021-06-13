package main

import (
	"testing"
)

/*
	## summary
	最大1文字削除後に回文になるか
*/
func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return helper(s, i, j-1) || helper(s, i+1, j)
		}
	}
	return true
}

func helper(str string, s, t int) bool {
	for s < t {
		if str[s] == str[t] {
			s++
			t--
		} else {
			return false
		}
	}
	return true
}

func TestValidPalindrome(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"aba", true,
		},
		{
			"abca", true,
		},
		{
			"abc", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := validPalindrome(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
