package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/*
	## summary
	find good tripletes
	arr[i], arr[j], arr[k] is good
	- i<j<k
	- |arr[i] - arr[j]| <= a ...
*/
func countGoodTriplets(arr []int, a int, b int, c int) int {
	l := len(arr)
	ret := 0
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			for k := j + 1; k < l; k++ {
				if (int(math.Abs(float64(arr[i]-arr[j]))) <= a) &&
					(int(math.Abs(float64(arr[j]-arr[k]))) <= b) &&
					(int(math.Abs(float64(arr[i]-arr[k]))) <= c) {
					ret++
				}
			}
		}
	}
	return ret
}

func TestCountGoodTriplets(t *testing.T) {
	tests := []struct {
		in      []int
		a, b, c int
		out     int
	}{
		{
			[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3, 4,
		},
		{
			[]int{1, 1, 2, 2, 3}, 0, 0, 1, 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countGoodTriplets(tt.in, tt.a, tt.b, tt.c)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
