package main

import (
	"strings"
	"testing"
)

/*
  ## summary
  seqの部分文字列wordの最大連携回数
*/
func maxRepeating(sequence string, word string) int {
	sl, wl := len(sequence), len(word)
	ret, tmp := 0, 0
	for i := 0; i <= sl-wl; i++ {
		if sequence[i:wl+i] == word {
			tmp++
			i += wl - 1
		} else {
			i -= tmp * wl
			tmp = 0
		}
		if ret < tmp {
			ret = tmp
		}
	}
	return ret
}

func maxRepeating2(sequence string, word string) int {
	ret := 1
	for strings.Contains(sequence, strings.Repeat(word, ret)) {
		ret++
	}
	return ret - 1
}

func TestIsMaxRepeating(t *testing.T) {
	tests := []struct {
		in1 string
		in2 string
		out int
	}{
		{
			"ababc", "ab", 2,
		},
		{
			"ababc", "ba", 1,
		},
		{
			"a", "a", 1,
		},
		{
			"aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba", 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in1+tt.in2, func(t *testing.T) {
			got := maxRepeating(tt.in1, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
