package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"
)

/*
	## summary
*/
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}

func arrayStringsAreEqual2(word1 []string, word2 []string) bool {
	s2 := strings.Join(word2, "")
	for _, w := range word1 {
		if !strings.HasPrefix(s2, w) {
			return false
		}
		s2 = s2[utf8.RuneCountInString(w):]
	}
	return len(s2) == 0
}

func TestArrayStringsAreEqual(t *testing.T) {
	tests := []struct {
		in  []string
		in2 []string
		out bool
	}{
		{
			[]string{"ab", "c"},
			[]string{"a", "bc"},
			true,
		},
		{
			[]string{"a", "cb"},
			[]string{"ab", "c"},
			false,
		},
		{
			[]string{"abc", "d", "defg"},
			[]string{"abcddefg"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := arrayStringsAreEqual2(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
