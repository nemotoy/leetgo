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
*/
func firstUniqChar(s string) int {

	size := len(s)
	for i := 0; i < size; i++ {
		f := false
		fmt.Printf("size: %d, s: %s\n", size, s[i:])
		// 基底値の次要素以降を対象
		for sindex, r := range s[i+1:] {
			fmt.Printf("bindex: %d, val: %s, sindex: %d, val: %s\n", i, string(s[i]), sindex, string(r))
			// 基底indexのruneと同値ならiterationを抜ける
			if r == rune(s[i]) {
				f = true
				fmt.Println(string(r), sindex, string(s[i]), f)
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
