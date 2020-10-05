package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	- 1文字ならindex 0 を返す
	- 先頭から1つずつ取り出して、全件検索
		- 同値がなければindexを保存
	- 全て見つからなければ-1を返す
*/
func oldfirstUniqChar(s string) int {

	size := len(s)
	// 1文字なら0を返す
	if size == 1 {
		return 0
	}
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

func firstUniqChar(s string) int {

	var count [26]int // アルファベット数分の長さをもつ配列を定義する

	for _, r := range s {
		// アルファベットのindexに要素1を加算する。各要素は文字列に含まれるアルファベットの数を示す。0は含まれない、1はユニーク。
		count[r-'a'] += 1
	}

	for i, r := range s {
		// 先頭から要素を検索し、1であるindexを返す。
		if count[r-'a'] == 1 {
			return i
		}
	}

	return -1
}

func TestFirstUniqChar(t *testing.T) {
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
