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
func reverse(x int) int {
	const MaxInt = int(math.MaxUint32 >> 1)
	const MinInt = -MaxInt - 1
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > MaxInt/10 || rev == MaxInt/10 && pop > 7 {
			return 0
		}
		if rev < MinInt/10 || rev == MinInt/10 && pop < -8 {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

func TestReverse(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			123, 321,
		},
		{
			-123, -321,
		},
		{
			120, 21,
		},
		{
			-2147483412, -2143847412,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := reverse(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
