package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	三連続奇数の場合は真を返す
*/
func threeConsecutiveOdds(arr []int) bool {
	c := 0
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			c++
		} else {
			c = 0
		}
		if c == 3 {
			return true
		}
	}
	return false
}

func TestThreeConsecutiveOdds(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{[]int{2, 6, 4, 1}, false},
		{[]int{1, 2, 34, 3, 4, 5, 7, 23, 1}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := threeConsecutiveOdds(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
