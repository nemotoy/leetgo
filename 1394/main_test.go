package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  配列内の出現頻度との値に等しい整数を返す。複数ある場合は大きい方を、ない場合は-1を返す
*/
func findLucky(arr []int) int {
	freq := make(map[int]int)
	for _, a := range arr {
		freq[a]++
	}

	ret := -1
	for k, v := range freq {
		if k == v && ret < k {
			ret = k
		}
	}

	return ret
}

func TestFindLucky(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{2, 2, 3, 4}, 2},
		{[]int{1, 2, 2, 3, 3, 3}, 3},
		{[]int{2, 2, 2, 3, 3}, -1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findLucky(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
