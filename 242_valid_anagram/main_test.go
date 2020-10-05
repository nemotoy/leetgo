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
func oldIsAnagram(s string, t string) bool {
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

func isAnagram(s string, t string) bool {
	var (
		count  [26]int
		count2 [26]int
	)

	for _, r := range s {
		count[r-'a'] += 1
	}
	for _, r := range t {
		count2[r-'a'] += 1
	}

	return reflect.DeepEqual(count, count2)
}

// func isAnagram(s string, t string) bool {
// 	var count [26]int

// 	if len(s) > len(t) {
// 		for _, r := range t {
// 			count[r-'a'] += 1
// 		}
// 		for _, r := range s {
// 			if count[r-'a'] -= 1; count[r-'a'] == -1 {
// 				return false
// 			}
// 		}
// 	}
// 	for _, r := range s {
// 		count[r-'a'] += 1
// 	}
// 	for _, r := range t {
// 		if count[r-'a'] -= 1; count[r-'a'] == -1 {
// 			return false
// 		}
// 	}
// 	return true
// }

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
		{
			"ab", "a", false,
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
