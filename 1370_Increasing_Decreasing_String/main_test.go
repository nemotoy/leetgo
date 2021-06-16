package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func sortString(s string) string {
	var chars [26]int
	for _, r := range s {
		chars[r-'a']++
	}
	ret, count, l := "", 0, len(s)
	for count < l {
		for i := 0; i < 26; i++ {
			if chars[i] > 0 {
				ret += string(rune(i + 'a'))
				chars[i]--
				count++
			}
		}
		for i := 25; i >= 0; i-- {
			if chars[i] > 0 {
				ret += string(rune(i + 'a'))
				chars[i]--
				count++
			}
		}
	}
	return ret
}

func TestSortString(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"aaaabbbbcccc", "abccbaabccba",
		},
		{
			"rat", "art",
		},
		{
			"leetcode", "cdelotee",
		},
		{
			"ggggggg", "ggggggg",
		},
		{
			"spo", "ops",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sortString(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
