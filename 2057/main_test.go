package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	i mod 10 == nums[i]が真である最小インデックスiを返す。ヒットしなければ-1を返す。
*/
func smallestEqual(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i%10 {
			return i
		}
	}
	return -1
}

func TestSmallestEqual(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{0, 1, 2}, 0},
		{[]int{4, 3, 2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := smallestEqual(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
