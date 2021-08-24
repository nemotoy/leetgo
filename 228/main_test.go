package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
	## summary
*/
func summaryRanges(nums []int) []string {
	l := len(nums)
	ret := make([]string, 0, len(nums))
	for i := 0; i < l; i++ {
		num := nums[i]
		for i < l-1 && nums[i]+1 == nums[i+1] {
			i++
		}
		if num != nums[i] {
			ret = append(ret, strconv.Itoa(num)+"->"+strconv.Itoa(nums[i]))
		} else {
			ret = append(ret, strconv.Itoa(num))
		}
	}
	return ret
}

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		in  []int
		out []string
	}{
		{
			[]int{0, 1, 2, 4, 5, 7},
			[]string{"0->2", "4->5", "7"},
		},
		{
			[]int{0, 2, 3, 4, 6, 8, 9},
			[]string{"0", "2->4", "6", "8->9"},
		},
		{
			[]int{},
			[]string{},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := summaryRanges(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
