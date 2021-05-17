package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func TestMySqrt(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{4, 2},
		{8, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := mySqrt(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
