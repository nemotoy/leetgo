package main

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

/*
	## summary
	numを2回reverseした結果とnumが一致するか。なお、12300のreverse結果は321とする。
*/
func isSameAfterReversals(num int) bool {
	return num == reverse(reverse(num))
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32/10 || rev == math.MaxInt32/10 && pop > 7 {
			return 0
		}
		if rev < math.MinInt32/10 || rev == math.MinInt32/10 && pop < -8 {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

func TestIsSameAfterReversals(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{
			526, true,
		},
		{
			1800, false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := isSameAfterReversals(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
