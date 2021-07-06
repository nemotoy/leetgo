package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func canBeEqual(target []int, arr []int) bool {
	bucket := [1001]int{}
	for i := 0; i < len(target); i++ {
		bucket[target[i]]++
		bucket[arr[i]]--
	}
	for _, v := range bucket {
		if v > 0 {
			return false
		}
	}
	return true
}

func TestCanBeEqual(t *testing.T) {
	tests := []struct {
		in  []int
		in2 []int
		out bool
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}, true},
		{[]int{7}, []int{7}, true},
		{[]int{1, 12}, []int{12, 1}, true},
		{[]int{3, 7, 9}, []int{3, 7, 11}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := canBeEqual(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
