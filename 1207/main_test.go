package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	Given an array of integers arr, write a function that returns true
	if and only if the number of occurrences of each value in the array is unique.
*/
func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	mm := make(map[int]int)

	for _, n := range arr {
		m[n] += 1
	}

	for _, v := range m {
		mm[v]++
		if mm[v] > 1 {
			return false
		}
	}
	return true
}

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := uniqueOccurrences(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
