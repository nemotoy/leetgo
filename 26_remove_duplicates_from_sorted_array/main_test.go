package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	配列の要素はソート済み。
	0番目をi, 1番目をjとし、各要素が同値でなければiをインクリメントし、j番目の要素をi番目に要素に入れる。
*/
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func Test_removeDuplicates(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 1, 2},
			2,
		},
		{
			[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			if got := removeDuplicates(tt.in); got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
