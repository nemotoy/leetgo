package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	sは大文字・小文字のアルファベットで構成される。
	- 0 <= i <= len(s)-2
	- s[i], s[i+1]が同じletterだが、大文字
	badな文字列はその要素同士を削除する。
*/
func makeGood(s string) string {
	l := len(s)
	for i := 0; i < l-1; {
		if IsSameLetter(s[i], s[i+1]) {
			s = s[:i] + s[i+2:]
			l -= 2
			if i > 0 {
				i--
			}
			continue
		}
		i++
	}
	return s
}

// 同じ文字か（ただし、大文字・小文字の組み合わせ）を評価する
func IsSameLetter(a, b byte) bool {
	return a == b+32 || a == b-32
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
