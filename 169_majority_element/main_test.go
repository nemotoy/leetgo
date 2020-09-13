package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
	サイズnの配列から最大値を探索する。
		- 最大値はn/2個以上含まれる
		- 空配列は与えられない

	- 配列の中の最大値を求める → 降順にする。
	- 配列サイズ/2の個数が含まれているかを見る → 降順後の配列をループして個数を数える
	- 最大値の個数 < n/2 である場合は対象外
*/
func majorityElement(nums []int) int {

	// 降順にする
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	size := len(nums)
	comparedItemNumber := float64(size) / float64(2)
	var count float64
	for i := 0; i < size; i++ {
		for _, num := range nums {
			if nums[i] == num {
				count++
			}
		}
		if count > comparedItemNumber {
			return nums[i]
		}
		count = 0
	}
	return 0
}

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{3, 2, 3},
			3,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
		{
			[]int{3, 3, 4},
			3,
		},
		{
			[]int{6, 6, 6, 7, 7},
			6,
		},
		{
			[]int{1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := majorityElement(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
