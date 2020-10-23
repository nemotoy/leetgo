package main

import (
	"strings"
	"testing"
)

/*
	## summary
*/
func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	res := ""
	l := len(ss) - 1
	for i, v := range ss {
		tail := " "
		if l == i {
			tail = ""
		}
		res += reverse(v) + tail
	}
	return res
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := reverseWords(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
