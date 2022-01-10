package main

import (
	"strconv"
	"strings"
	"testing"
	"unicode"
)

/*
	## summary
	sに含まれる整数が厳密な昇順
*/
func areNumbersAscending(s string) bool {
	ss := strings.Split(s, " ")
	tmp := 0
	for _, v := range ss {
		first := rune(v[0])
		if unicode.IsDigit(first) {
			num, _ := strconv.Atoi(v)
			if tmp < num {
				tmp = num
			} else {
				return false
			}
		}
	}
	return true
}

func TestAreNumbersAscending(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"1 box has 3 blue 4 red 6 green and 12 yellow marbles", true,
		},
		{
			"hello world 5 x 5", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := areNumbersAscending(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
