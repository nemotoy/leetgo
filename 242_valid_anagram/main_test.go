package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	- 長さが異なる場合はfalse
	- 文字列型を1runeずつ配列の要素に追加する
	- 配列をalphebetiacalにソートする
	- 配列が同値か比較する
*/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ss := make([]string, len(s))
	tt := make([]string, len(t))
	for _, v := range s {
		ss = append(ss, string(v))
	}
	for _, v := range t {
		tt = append(tt, string(v))
	}

	sort.Strings(ss)
	sort.Strings(tt)
	return reflect.DeepEqual(ss, tt)
}

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		inS string
		inT string
		out bool
	}{
		{
			"anagram", "nagaram", true,
		},
		{
			"rat", "car", false,
		},
		{
			"", "", true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %s", tt.inS, tt.inT), func(t *testing.T) {
			got := isAnagram(tt.inS, tt.inT)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
