package main

import (
	"fmt"
	"math/bits"
	"reflect"
	"testing"
)

/*
	## summary
	2進数で表現した際の差異。0 ≤ x, y < 231
	文字数 n の1ビット文字列間のハミング距離は、それらの文字列間の排他的論理和のハミング重み（文字列内の 1 の個数）に相当する。
*/
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		in1 int
		in2 int
		out int
	}{
		{
			1, 4, 2,
		},
		{
			3, 1, 1,
		},
		{
			93, 73, 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d", tt.in1, tt.in2), func(t *testing.T) {
			got := hammingDistance(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
