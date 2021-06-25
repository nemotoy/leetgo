package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
	wordsの各要素においてallowedの文字列で構成されている数
*/
func countConsistentStrings(allowed string, words []string) int {
	ret := 0
	for _, word := range words {
		increment := true
		for _, r := range word {
			if !strings.Contains(allowed, string(r)) {
				increment = false
				break
			}
		}
		if increment {
			ret++
		}
	}
	return ret
}

func TestCountConsistentStrings(t *testing.T) {
	tests := []struct {
		in  string
		in2 []string
		out int
	}{
		{
			"ab", []string{"ad", "bd", "aaab", "baa", "badab"}, 2,
		},
		{
			"abc", []string{"a", "b", "c", "ab", "ac", "bc", "abc"}, 7,
		},
		{
			"cad", []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}, 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countConsistentStrings(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
