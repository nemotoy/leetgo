package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## memo
	- 0の個数をカウントする
	- 個数分の0を含むリストを作る
	- 与えられたリストから要素0を削除する（ ※ incrementと要素がズレる）
	- リストとリストを結合する
*/
func moveZeroes(nums []int) {
	incList := []int{}
	for i, n := range nums {
		fmt.Printf("#1 inc: %d, nums: %v\n", i, nums)
		if n == 0 {
			incList = append(incList, i)
		}
	}
	count := len(incList)
	fmt.Printf("#1 count: %d\n", len(incList))
	zeroList := make([]int, count)
	for count < 0 {
		zeroList = append(zeroList, 0)
		count--
	}
	fmt.Printf("#2 zero list: %v\n", zeroList)

	base := 0
	for _, n := range incList {
		fmt.Printf("#3 n: %d, nums: %v\n", n, nums)
		n = n - base
		if nums[n] == 0 {
			fmt.Printf("Prev; %v, %v\n", nums[:n], nums[n+1:])
			nums = append(nums[:n], nums[n+1:]...)
			fmt.Printf("follow; nums: %v\n", nums)
			base++
		}
	}
	nums = append(nums, zeroList...)
	fmt.Printf("#4 nums: %v\n", nums)
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
