package main

import (
	"fmt"
	"testing"
)

/*
  ## summary
  'balloon'が含まれる個数
*/
func maxNumberOfBalloons(text string) int {
	if len(text) < 7 {
		return 0
	}
	ret := 0
	for len(text) >= 7 {
		for _, w := range "balloon" {
			prev := len(text)
			for it, t := range text {
				if w == t {
					text = text[:it] + text[it+1:]
					break
				}
			}
			if prev-1 != len(text) {
				return ret
			}
		}
		ret++
	}
	return ret
}

func maxNumberOfBalloons2(text string) int {
	// アルファベット数の要素を持つ配列を初期化する
	var chars [26]int
	// inputに含まれるアルファベット数を集計する
	for _, r := range text {
		chars[r-'a'] += 1
	}
	ret := chars[1]             // b
	ret = min(ret, chars[0])    // a
	ret = min(ret, chars[11]/2) // l*2
	ret = min(ret, chars[14]/2) // o*2
	ret = min(ret, chars[13])   // n
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMaxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"nlaebolko", 1,
		},
		{
			"loonbalxballpoon", 2,
		},
		{
			"leetcode", 0,
		},
		{
			"balloon", 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxNumberOfBalloons2(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
