package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func findMiddleIndex(nums []int) int {
	for i := 0; i < len(nums); i++ {
		left, right := nums[:i], nums[i+1:]
		if sum(left) == sum(right) {
			return i
		}
	}
	return -1
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func TestFindMiddleIndex(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{2, 3, -1, 8, 4}, 3},
		{[]int{1, -1, 4}, 2},
		{[]int{2, 5}, -1},
		{[]int{0, 0, 0, 0}, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findMiddleIndex(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
