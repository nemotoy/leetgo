package main

import (
	"testing"
)

/*
	## summary
	word1, word2の順で文字列を結合した結果を返す。
*/
func mergeAlternately(word1 string, word2 string) string {
	ret := ""
	l1, l2 := len(word1), len(word2)
	for i := 0; i < l1 || i < l2; i++ {
		if i < l1 {
			ret += string(word1[i])
		}
		if i < l2 {
			ret += string(word2[i])
		}
	}
	return ret
}

func TestMergeAlternately(t *testing.T) {
	tests := []struct {
		in  string
		in2 string
		out string
	}{
		{
			"abc", "pqr", "apbqcr",
		},
		{
			"ab", "pqrs", "apbqrs",
		},
		{
			"abcd", "pq", "apbqcd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := mergeAlternately(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
