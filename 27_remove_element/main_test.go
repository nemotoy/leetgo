package main

import (
	"fmt"
	"reflect"
	"testing"
)

// valと等価でない値の個数kを返し、numsは先頭からk個分valと等価でない要素が含まれるようにする
func removeElement(nums []int, val int) int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums:%v target:%d", tt.in, tt.in2), func(t *testing.T) {
			got := removeElement(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
