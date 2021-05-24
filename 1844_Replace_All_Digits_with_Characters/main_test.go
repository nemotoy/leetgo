package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
  ## summary
  alphabets, number...の順に並べられた文字列から、数値を左隣のアルファベットを数値分進めた値にして返す。
*/
func replaceDigits(s string) string {
	l, i := len(s)-1, 0
	for i < l {
		n, _ := strconv.Atoi(string(s[i+1]))
		r := rune(s[i])
		tmp := shiftAlps(r, n)
		s = s[:i+1] + string(tmp) + s[i+2:]
		i += 2
	}
	return s
}

func shiftAlps(r rune, n int) rune {
	alps := "abcdefghijklmnopqrstuvwxyz"
	for i, a := range alps {
		if r == a {
			width := i + n
			if width >= 26 {
				return rune(byte('z'))
			}
			return rune(alps[width])
		}
	}
	return 0
}

func replaceDigits2(s string) string {
	ret := []byte(s)
	// 奇数番目を走査する。
	for i := 1; i < len(s); i += 2 {
		ret[i] = ret[i-1] + ret[i] - '0'
	}
	return string(ret)
}

func TestReplaceDigits(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"a1c1e1", "abcdef",
		},
		{
			"a1b2c3d4e", "abbdcfdhe",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := replaceDigits2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
