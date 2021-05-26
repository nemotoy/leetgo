package main

import (
	"testing"
)

/*
  ## summary
  sは偶数長。両文字列の母音数が同数であるかを評価する。なお、母音数は大文字・小文字を含む。
*/
func halvesAreAlike(s string) bool {
	l := len(s) / 2
	a, b := s[:l], s[l:]
	return getVowelsNumber(a) == getVowelsNumber(b)
}

func getVowelsNumber(s string) int {
	ret := 0
	// strings.ToLower(s), 'a', 'e', 'i', 'o', 'u'でも書けるが、メモリ効率が悪い
	for _, r := range s {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			ret++
		}
	}
	return ret
}

func TestHalvesAreAlike(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"book", true,
		},
		{
			"textbook", false,
		},
		{
			"MerryChristmas", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := halvesAreAlike(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
