package main

import (
	"fmt"
	"reflect"
	"testing"
)

// https://brilliant.org/wiki/trailing-number-of-zeros/
func trailingZeroes(n int) int {
	i, res := 5, 0
	for n >= i {
		res += n / i
		i *= 5
	}
	return res
}

func TestTrailingZeroes(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{3, 0},
		{5, 1},
		{0, 0},
		{7, 1},
		{4, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := trailingZeroes(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
