package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	ith
*/
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ret := 0
	for i := 0; i < len(startTime); i++ {
		// startTime[i] <= queryTime <= endTime[i]
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			ret++
		}
	}
	return ret
}

func TestBusyStudent(t *testing.T) {
	tests := []struct {
		in  []int
		in2 []int
		in3 int
		out int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
		{[]int{4}, []int{4}, 4, 1},
		{[]int{4}, []int{4}, 5, 0},
		{[]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7, 0},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := busyStudent(tt.in, tt.in2, tt.in3)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
