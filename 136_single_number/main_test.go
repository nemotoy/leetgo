package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	空でないint型の配列から、ユニークな要素を取得する。
	- 重複した値は3回以上エントリされない
	- 重複しない値は必ず1つである

	1. Hash table
	- 配列の全ての要素を探索して、ハッシュに値を追加する。その際、重複があれば削除する。これが利用できるのは、前提条件があるから。
	- 最終的なハッシュテーブルに残る要素は前提条件より必ず1つになるため、それを返す。

	2. Bit Manipulation
	- a XOR 0 = a
	- a XOR a = 0
*/
func singleNumber(nums []int) int {

	table := make(map[int]int)
	for _, num := range nums {
		if vt, ok := table[num]; !ok { // 重複なし
			table[num] = 1
		} else if vt == 1 { // 重複あり
			delete(table, num)
		}
	}
	// テーブルの要素（長さ）は必ず 1
	for v := range table {
		return v
	}
	return 0
}

func singleNumber2(nums []int) int {
	a := 0
	for _, n := range nums {
		a ^= n
	}
	return a
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{2, 2, 1}, 1,
		},
		{
			[]int{4, 1, 2, 1, 2}, 4,
		},
		{
			[]int{1, 3, 1, -1, 3},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := singleNumber(tt.in)
			if got != tt.out {
				t.Errorf("got: %d, want: %d", got, tt.out)
			}
		})
	}
}
