package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	[freq, val] = [nums[2*i], nums[2*i+1]]
	freq個分だけ要素valを追加
*/
func decompressRLElist(nums []int) []int {
	ret := []int{}
	for i := 0; i < len(nums); {
		freq, val := nums[i], nums[i+1]
		vals := make([]int, freq)
		for j := 0; j < freq; j++ {
			vals[j] = val
		}
		ret = append(ret, vals...)
		i += 2
	}
	return ret
}

func TestDecompressRLElist(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{2, 4, 4, 4},
		},
		{
			[]int{1, 1, 2, 3},
			[]int{1, 3, 3},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := decompressRLElist(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
