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
		- 最大値はn/2回含まれる
		- 空配列は与えられない

	- 配列の中の最大値を求める → 降順にする。
	- 配列サイズ/2の個数が含まれているかを見る → 降順後の配列をループして個数を数える
*/
func majorityElement(nums []int) int {

	// 降順
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums)
	var size = len(nums) / 2
	fmt.Println("size: ", size)

	count := []int{}
	for i, n := range nums {
		fmt.Println("inc: ", i)
		prev := i - 1
		// 0番目の場合
		if prev == -1 {
			count = append(count, n)
			continue
		}
		// 前要素と同値の場合
		if n == nums[prev] {
			count = append(count, n)
		} else {
			// 同値でない場合
			fmt.Println("count: , size: ", len(count), size)
			// sizeより大きければ返す
			if len(count) > size {
				fmt.Println("val: ", count)
				return count[0]
			}
			// それ以外はその要素でリセットする
			count = []int{}
			count = append(count, n)
			fmt.Println("val: ", count)
		}
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
