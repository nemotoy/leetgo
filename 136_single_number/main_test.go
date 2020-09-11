package main

import (
	"fmt"
	"testing"
)

func singleNumber(nums []int) int {

	dup := []int{}
	table := make(map[int]int)
	for _, v := range nums {
		if _, ok := table[v]; ok {
			dup = append(dup, v)
			continue
		}
		table[v] = v
	}
	for _, v := range dup {
		if _, ok := table[v]; ok {
			delete(table, v)
		}
	}
	for v := range table {
		return v
	}
	return 0
}

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{2, 2, 1}, 1,
		},
		{
			[]int{4, 1, 2, 1, 2}, 4,
		},
		{
			[]int{1, 3, 1, -1, 3},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := singleNumber(tt.in)
			if got != tt.out {
				t.Errorf("got: %d, want: %d", got, tt.out)
			}
		})
	}
}
