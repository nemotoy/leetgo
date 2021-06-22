package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	奇数長の合計値
*/
func sumOddLengthSubarrays(arr []int) int {
	ret, l := 0, len(arr)
	for i, a := range arr {
		ret += ((i+1)*(l-i) + 1) / 2 * a
	}
	return ret
}

func TestSumOddLengthSubarrays(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 4, 2, 5, 3},
			58,
		},
		{
			[]int{1, 2},
			3,
		},
		{
			[]int{10, 11, 12},
			66,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sumOddLengthSubarrays(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
