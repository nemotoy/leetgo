package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	?をlower case lettersに置換する。それは連続した文字列にならない。
	1つの?の解は24or25パターンある。
*/
func modifyString(s string) string {
	// Strings are immutable: once created, it is impossible to change the contents of a string
	// https://golang.org/ref/spec#String_types
	// string型はimmutableなので、runeの配列からstring型に変換する
	b := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			for _, c := range "abc" {
				if (i == 0 || b[i-1] != c) && (i+1 == len(s) || b[i+1] != c) {
					b[i] = c
					break
				}
			}
		}
	}
	return string(b)
}

func TestModifyString(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"?zs", "azs",
		},
		{
			"ubv?w", "ubvaw",
		},
		{
			"j?qg??b", "jaqgacb",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := modifyString(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
