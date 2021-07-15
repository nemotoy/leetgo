package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func findPoisonedDuration(timeSeries []int, duration int) int {
	l := len(timeSeries)
	if l == 0 {
		return 0
	}
	ret := 0
	for i := 0; i < l-1; i++ {
		ret += min(timeSeries[i+1]-timeSeries[i], duration)
	}
	return ret + duration
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestFindPoisonedDuration(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{
			[]int{1, 4}, 2, 4,
		},
		{
			[]int{1, 2}, 2, 3,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums:%v target:%d", tt.in, tt.in2), func(t *testing.T) {
			got := findPoisonedDuration(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
