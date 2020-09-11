package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	- 1度でも重複したらtrue
*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for i, num := range nums {
		if i == 0 {
			m[num] = struct{}{}
			continue
		}
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
