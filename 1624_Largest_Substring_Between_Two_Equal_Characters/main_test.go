package main

import (
	"strings"
	"testing"
)

/*
	## summary
	同値要素間のindex差分。
*/
func maxLengthBetweenEqualCharacters(s string) int {
	res := -1
	for _, r := range s {
		firstIndx, lastIndx := strings.Index(s, string(r)), strings.LastIndex(s, string(r))
		// strings.Index/strings.LastIndex は該当しない場合 -1を返すので、基底値-1より小さくなり代入されることはない。
		res = Max(res, lastIndx-firstIndx-1)
	}
	return res
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestMaxLengthBetweenEqualCharacters(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"aa", 0,
		},
		{
			"abca", 2,
		},
		{
			"cbzxy", -1,
		},
		{
			"ojdncpvyneq", 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := maxLengthBetweenEqualCharacters(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
