package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	nums.Length == 2n

	配列をn以前・以降の2つに分けて、インデックスが交互になるようにする。
*/
func shuffle(nums []int, n int) []int {
	if n == 1 {
		return nums
	}
	xv, yv := nums[:n], nums[n:]
	fmt.Println(xv, yv)
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		switch {
		case i%2 == 0:
			result[i] = xv[i]
		default:
			result[i] = yv[i]
		}
	}
	return result
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		in  []int
		inN int
		out []int
	}{
		{
			[]int{2, 5, 1, 3, 4, 7},
			3,
			[]int{2, 3, 5, 4, 1, 7},
		},
		{
			[]int{1, 2, 3, 4, 4, 3, 2, 1},
			4,
			[]int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			[]int{1, 1, 2, 2},
			2,
			[]int{1, 2, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.inN), func(t *testing.T) {
			got := shuffle(tt.in, tt.inN)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
