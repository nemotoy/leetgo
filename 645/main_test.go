package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	Find the number that occurs twice and the number that is missing and return them in the form of an array.
*/
func findErrorNums(nums []int) []int {
	numMap := make(map[int]int, len(nums))
	dup, missing := -1, -1
	for _, num := range nums {
		numMap[num] += 1
	}
	// 本来のnumsは1~nまでの重複しない整数が含まれるので、n = len(nums)が成立する
	for i := 1; i <= len(nums); i++ {
		v, ok := numMap[i]
		if ok {
			if v == 2 {
				dup = i
			}
		} else {
			missing = i
		}
	}
	return []int{dup, missing}
}

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{1, 2, 2, 4},
			[]int{2, 3},
		},
		{
			[]int{1, 1},
			[]int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findErrorNums(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
