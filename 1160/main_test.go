package main

import (
	"fmt"
	"testing"
)

/*
  ## summary
  charsで構成されるwords[i]の長さの合計値
*/
func countCharacters(words []string, chars string) int {
	var count [26]int
	for _, r := range chars {
		count[r-'a'] += 1
	}
	ret := 0
	for _, word := range words {
		match := true
		tmp := count
		for _, r := range word {
			tmp[r-'a'] -= 1
			if tmp[r-'a'] < 0 {
				match = false
				break
			}
		}
		if match {
			ret += len(word)
		}
	}
	return ret
}

func TestCountCharacters(t *testing.T) {
	tests := []struct {
		in  []string
		in2 string
		out int
	}{
		{
			[]string{"cat", "bt", "hat", "tree"}, "atach", 6,
		},
		{
			[]string{"hello", "world", "leetcode"}, "welldonehoneyr", 10,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %s", tt.in, tt.in2), func(t *testing.T) {
			got := countCharacters(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
