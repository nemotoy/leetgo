package main

import (
	"bytes"
	"strings"
	"testing"
)

/*
	## summary
*/
func reverseWords(s string) string {

	words := strings.Split(s, " ")
	sb := &bytes.Buffer{}
	l := len(words)
	for i, s := range words {
		var sr = []byte(s)
		reverseString(sr)
		sb.Write(sr)
		if i < l-1 {
			sb.WriteRune(' ')
		}
	}
	return sb.String()
}

func reverseString(s []byte) {
	for i, j := len(s)-1, 0; i > j; i, j = i-1, j+1 {
		s[i], s[j] = s[j], s[i]
	}
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
