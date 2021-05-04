package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/*
	## summary
*/
func shortestToChar(s string, c byte) []int {
	l := len(s)
	result := make([]int, l)
	prev := math.MinInt32 / 2
	for i := 0; i < l; i++ {
		if s[i] == c {
			prev = i
		}
		result[i] = i - prev
	}
	prev = math.MaxInt32 / 2
	for i := l - 1; i >= 0; i-- {
		if s[i] == c {
			prev = i
		}
		result[i] = min(result[i], prev-i)
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestShortestToChar(t *testing.T) {
	tests := []struct {
		in  string
		in2 byte
		out []int
	}{
		{
			"loveleetcode", byte('e'), []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			"aaab", byte('b'), []int{3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.in2), func(t *testing.T) {
			got := shortestToChar(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
