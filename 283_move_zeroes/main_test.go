package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	- 0の個数をカウントする
	- 個数分の0を含むリストを作る
	- 与えられたリストから要素0を削除する
	- リスト同士を結合する
*/
func moveZeroes(nums []int) {
	incList := []int{}
	for i, n := range nums {
		if n == 0 {
			incList = append(incList, i)
		}
	}
	count := len(incList)
	zeroList := make([]int, count)
	for count < 0 {
		zeroList = append(zeroList, 0)
		count--
	}

	base := 0
	for _, n := range incList {
		// インクリメントをリフレッシュする
		n = n - base
		if nums[n] == 0 {
			// 対象の要素を削除する
			nums = append(nums[:n], nums[n+1:]...)
			base++
		}
	}
	nums = append(nums, zeroList...)
}

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{0, 1, 0, 3, 12},
			[]int{1, 3, 12, 0, 0},
		},
		{
			[]int{0, 0, 1},
			[]int{1, 0, 0},
		},
		{
			[]int{1, 0},
			[]int{1, 0},
		},
		{
			[]int{1, 0, 0},
			[]int{1, 0, 0},
		},
		{
			[]int{1, 0, 1},
			[]int{1, 1, 0},
		},
		{
			[]int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0},
			[]int{4, 2, 4, 3, 5, 1, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			moveZeroes(tt.in)
			got := tt.in
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
