package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	先頭要素を基底にして、他配列要素を走査する。
*/
func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return strs[0]
	}
	base := strs[0]
	prefix := ""
	for i := 1; i <= len(base); i++ {
		// 配列の1要素目をprefixとし、先頭から1文字ずつ走査する
		prefix = base[:i]
		// 配列の2要素目以降を走査する
		for _, s := range strs[1:] {
			// sの先頭からprefixの長さ分の文字列と比較する。先に長さの比較を行う。
			if len(s) < len(prefix) || s[:len(prefix)] != prefix {
				// 前回走査したprefixを返す
				return base[:i-1]
			}
		}
	}
	return prefix
}

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		in  []string
		out string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			[]string{"dog", "racecar", "car"},
			"",
		},
		{
			[]string{"a"},
			"a",
		},
		{
			[]string{"c", "c"},
			"c",
		},
		{
			[]string{"a", "b"},
			"",
		},
		{
			[]string{"c", "acc", "ccc"},
			"",
		},
		{
			[]string{"aa", "a"},
			"a",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := longestCommonPrefix(tt.in)
			if got != tt.out {
				t.Errorf("got: %s, want: %s", got, tt.out)
			}
		})
	}
}
