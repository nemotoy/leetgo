package main

import (
	"fmt"
	"reflect"
	"testing"
	"unicode"
)

/*
	## summary

	1. All letters in this word are capitals, like "USA".
	2. All letters in this word are not capitals, like "leetcode".
	3. Only the first letter in this word is capital, like "Google".

	upper letter数を計測して、上記の条件式を定義する。
*/
func detectCapitalUse(word string) bool {
	var c int
	for _, r := range word {
		if unicode.IsUpper(r) {
			c++
		}
	}
	return c == len(word) || c == 0 || (c == 1 && unicode.IsUpper(rune(word[0])))
}

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"USA", true,
		},
		{
			"FlaG", false,
		},
		{
			"ffffffffffffffffffffF", false,
		},
		{
			"Leetcode", true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := detectCapitalUse(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
