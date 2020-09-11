package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	- 重複した値は3回以上エントリされない
	- 重複しない値は必ず1つである
*/
func singleNumber(nums []int) int {

	table := make(map[int]int)
	for _, nums := range nums {
		if vt, ok := table[nums]; !ok { // 重複なし
			table[nums] = 1
		} else if vt == 1 { // 重複あり
			delete(table, nums)
		}
	}
	// テーブルの要素（長さ）は必ず 1
	for v := range table {
		return v
	}
	return 0
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
