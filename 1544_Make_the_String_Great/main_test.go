package main

import (
	"fmt"
	"testing"
	"unicode"
)

/*
	## summary
	sは大文字・小文字のアルファベットで構成される。
	- 0 <= i <= len(s)-2
	- s[i], s[i+1]が同じletterだが、大文字
	badな文字列はその要素同士を削除する。
*/
func makeGood(s string) string {
	if len(s) < 2 {
		return s
	}
	i, l := 0, len(s)-1
	for i < l {
		if unicode.IsUpper(rune(s[i])) {
			fmt.Println("upper", i, s[i], s[i+1], string(s[i]), string(s[i+1]))
			if IsSameLetter(s[i+1], s[i]) {
				if len(s) == 2 {
					s = ""
					break
				}
				s = s[:i] + s[i+2:]
				i, l = 0, len(s)-1
				fmt.Println("upper.1", s, i, l)
			}
		} else {
			if IsSameLetter(s[i], s[i+1]) {
				fmt.Println("lower")
				if len(s) == 2 {
					s = ""
					break
				}
				s = s[:i] + s[i+2:]
				i, l = 0, len(s)-1
				fmt.Println("upper.1", s, i, l)
			}
		}
		i++
	}
	return s
}

// a > b
func IsSameLetter(a, b byte) bool {
	return a > b && a-b == 32
}

func TestMakeGood(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"leEeetcode", "leetcode",
		},
		{
			"abBAcC", "",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := makeGood(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
