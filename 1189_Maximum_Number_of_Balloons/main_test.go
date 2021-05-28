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
			got := maxNumberOfBalloons(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
