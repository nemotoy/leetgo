package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
	## summary
	正の整数の各要素の二乗の合計が1になる場合、trueを返す。
*/
func isHappy(n int) bool {
	amount := 0
	if n < 1 {
		return false
	}
	s := strconv.Itoa(n)
	for _, r := range s {
		nn, _ := strconv.Atoi(string(r))
		amount += nn * nn
		if amount == 1 {
			return true
		}
	}
	return false
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{
			19, true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := isHappy(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
