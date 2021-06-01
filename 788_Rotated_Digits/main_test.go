package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  ## summary
  ローテート後、異なる値ならば正。
*/
func rotatedDigits(n int) int {
	ret := 0
	for i := 0; i <= n; i++ {
		if helperRotatedDigits(i) {
			ret++
		}
	}
	return ret
}

func helperRotatedDigits(n int) bool {
	ret := false
	for n > 0 {
		switch n % 10 {
		case 2, 5, 6, 9:
			ret = true
		case 3, 4, 7:
			return false
		}
		n /= 10
	}
	return ret
}

func TestRotatedDigits(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{10, 4},
		{2, 1},
		{857, 247},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := rotatedDigits(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
