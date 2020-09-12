package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	- 先頭から1つずつ取り出して、全件検索
		- 同値がなければindexを保存
	- 全て見つからなければ-1を返す
	- 1文字ならindex 0 を返す
*/
func firstUniqChar(s string) int {

	size := len(s)
	// 1文字なら0を返す
	if size == 1 {
		return 0
	}
	// 最後尾は処理対象外
	for i := 0; i < size; i++ {
		f := false
		// indexを除いた文字列を作る
		ss := s[:i] + s[i+1:]
		for _, r := range ss {
			// 基底indexのruneと同値ならiterationを抜ける
			if r == rune(s[i]) {
				f = true
				break
			}
		}
		// 同値がない場合は、indexを返す
		if !f {
			return i
		}
	}
	return -1
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"leetcode",
			0,
		},
		{
			"loveleetcode",
			2,
		},
		{
			"cc",
			-1,
		},
		{
			"z",
			0,
		},
		{
			"aadadaad",
			-1,
		},
		{
			"dddccdbba",
			8,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := firstUniqChar(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
