package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	arr.len >= 3
	arr[0] < arr[1] < ... arr[i-1] < arr[i]
	arr[i] > arr[i+1] > ... arr[arr.len-1]

	i-1, i+1が条件なので0とarr.lenは該当しない
*/
func peakIndexInMountainArray(arr []int) int {
	ret := 0
	for i := 1; i < len(arr)-1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			ret = i
			break
		}
	}
	return ret
}

func TestPeakIndexInMountainArray(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{0, 1, 0}, 1},
		{[]int{0, 2, 1, 0}, 1},
		{[]int{0, 10, 5, 2}, 1},
		{[]int{3, 4, 5, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := peakIndexInMountainArray(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
