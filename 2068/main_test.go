package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	両文字列に含まれるa-zまでの文字数の差分が3以下の場合、真。
*/
func checkAlmostEquivalent(word1 string, word2 string) bool {
	cnt := [2][27]int{{}}
	for _, w := range word1 {
		cnt[0][w-'a']++
	}
	for _, w := range word2 {
		cnt[1][w-'a']++
	}
	for i := 0; i < 27; i++ {
		if abs(cnt[0][i]-cnt[1][i]) > 3 {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x < 1 {
		return -x
	}
	return x
}

func TestCheckAlmostEquivalent(t *testing.T) {
	tests := []struct {
		in  string
		in2 string
		out bool
	}{
		{
			"aaaa", "bccb", false,
		},
		{
			"abcdeef", "abaaacc", true,
		},
		{
			"zzzyyy", "iiiiii", false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %s", tt.in, tt.in2), func(t *testing.T) {
			got := checkAlmostEquivalent(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
