package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	int型リストの0~nの中で存在しない要素を探索する
	- 最大値 n を決める
	- 0からnまでのリストあるいはマップを作る
	- 与えられたint型リストをイテレーションし、存在するか検索する
*/
func missingNumber(nums []int) int {

	if len(nums) == 1 && nums[0] == 0 {
		return 1
	}

	max := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if max < v {
			max = v
		}
	}

	m := make(map[int]struct{}, max)
	for i := 0; i <= max; i++ {
		m[i] = struct{}{}
	}

	for _, num := range nums {
		if _, ok := m[num]; ok {
			delete(m, num)
		}
	}

	for v := range m {
		return v
	}

	return 0
}

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{3, 0, 1},
			2,
		},
		{
			[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			8,
		},
		{
			[]int{0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := missingNumber(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
