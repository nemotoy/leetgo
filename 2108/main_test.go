package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	最初にヒットしたpalindromicな文字列を返す。
*/
func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}
	return ""
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func TestFirstPalindrome(t *testing.T) {
	tests := []struct {
		in  []string
		out string
	}{
		{
			[]string{"abc", "car", "ada", "racecar", "cool"},
			"ada",
		},
		{
			[]string{"notapalindrome", "racecar"},
			"racecar",
		},
		{
			[]string{"xngla", "e", "itrn", "y", "s", "pfp", "rfd"},
			"e",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := firstPalindrome(tt.in)
			if got != tt.out {
				t.Errorf("got: %s, want: %s", got, tt.out)
			}
		})
	}
}
