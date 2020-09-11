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
		for sindex, r := range s {
			// 基底値の次index以降を対象
			if i < sindex {
				fmt.Println(string(r), sindex, string(s[i]))
				// 同値ならiterationを抜ける
				if r == rune(s[i]) {
					f = true
					break
				}
			}
		}
		// フラグがfalseの場合（同値の値がない）は、indexを返す
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
