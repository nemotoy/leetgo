package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	- 1度でも重複したらtrue

	- sorting
	ソートして、隣り合う要素が同値か評価し続ける。
	- hash table

*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{
			[]int{1, 2, 3, 1},
			true,
		},
		{
			[]int{1, 2, 3, 4},
			false,
		},
		{
			[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := containsDuplicate(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
