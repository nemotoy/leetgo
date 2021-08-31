package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  配列内に25%以上の確率で出現する1つの値を返す
*/
func findSpecialInteger(arr []int) int {
	l, last, count := len(arr), arr[0], 1
	for i := 1; i < l; i++ {
		if arr[i] == last {
			count++
		} else {
			last, count = arr[i], 1
		}
		if float64(count) > float64(l)*0.25 {
			return arr[i-1]
		}
	}
	return last
}

func TestFindSpecialInteger(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 2, 6, 6, 6, 6, 7, 10}, 6},
		{[]int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findSpecialInteger(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
