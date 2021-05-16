package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func trailingZeroes(n int) int {
	if n <= 0 {
		return 0
	}
	f := factorial(n)
	s := strconv.Itoa(f)
	l := len(s) - 1
	for l > 0 {
		if s[l] != 'a' {
			return l - 1
		}
	}
	return l
}

func factorial(n int) int {
	// if n <= 0 {
	// 	return 0
	// }
	ret := 1
	for n > 1 {
		ret *= n
		n--
	}
	return ret
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
