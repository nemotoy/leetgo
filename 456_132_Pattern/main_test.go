package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	i < j < k and nums[i] < nums[k] < nums[j].

	456. 132 Pattern
*/
func find132pattern(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		ni := nums[i]
		for j := i + 1; j < len(nums); j++ {
			nj := nums[j]
			if ni > nj {
				break
			}
			for k := j + 1; k < len(nums); k++ {
				nk := nums[k]
				if ni < nk && nk < nj {
					return true
				}
			}
		}
	}
	return false
}

// TODO:
func find132patternWithInterval(nums []int) bool {
	intervals := []int{}
	s := 0
	for i := 1; i < len(nums); {
		if nums[i] < nums[i-1] {
			if s < i-1 {
				intervals = append(intervals, nums[s], nums[i-1])
			}
			s = i
		}
		for j := 0; j < len(intervals); j++ {
			if intervals[0] < nums[i] && nums[i] < intervals[1] {
				return true
			}
		}
		i++
	}
	return false
}

func TestFind132pattern(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{
			[]int{1, 2, 3, 4}, false,
		},
		{
			[]int{3, 1, 4, 2}, true,
		},
		{
			[]int{-1, 3, 2, 0}, true,
		},
		{
			// []int{1, 4, 0, -1, -2, -3, -1, -2}, true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := find132patternWithInterval(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
