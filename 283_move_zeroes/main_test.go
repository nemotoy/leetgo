package main

import (
	"fmt"
	"reflect"
	"testing"
)

func moveZeroes(nums []int) {
	incList := []int{}
	for i, n := range nums {
		fmt.Printf("#1 inc: %d, nums: %v\n", i, nums)
		if n == 0 {
			incList = append(incList, i)
		}
	}
	fmt.Printf("#2 increment list: %v, nums: %v\n", incList, nums)
	for _, inc := range incList {
		fmt.Printf("#3 nums: %v, length: %d, inc: %d\n", nums, len(nums), inc)
		if len(nums)-1 == inc {
			continue
		}
		if inc != 0 {
			inc--
		}
		if nums[inc] == 0 {
			nums = append(nums[:inc], nums[inc+1:]...)
			nums = append(nums, 0)
			fmt.Printf("#3.5 nums: %v\n", nums)
		}
	}
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
