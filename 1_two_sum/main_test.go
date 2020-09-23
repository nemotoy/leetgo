package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	numsの要素を2つ加算し、targetになるインデックス群を返す。

	Sol1. ブルートフォース
	基底値をi,次基底値をj(i+1)とし、i+j(jはlen(nums)まで検索) = targetを順次比較する。

	Sol2. ハッシュテーブル
*/
func twoSum(nums []int, target int) []int {
	for i, n1 := range nums {
		for j, n2 := range nums[i+1:] {
			if target == n1+n2 {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{}
}

func twoSumWithHashTable(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if v, ok := m[complement]; ok {
			if v != i {
				return []int{i, v}
			}
		}
	}
	return []int{}
}

func Test(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out []int
	}{
		{
			[]int{2, 7, 11, 15},
			9,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 4},
			6,
			[]int{1, 2},
		},
		{
			[]int{3, 3},
			6,
			[]int{0, 1},
		},
		{
			[]int{3, 2, 3},
			6,
			[]int{0, 2},
		},
		{
			[]int{2, 5, 5, 11},
			10,
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums:%v target:%d", tt.in, tt.in2), func(t *testing.T) {
			got := twoSumWithHashTable(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
