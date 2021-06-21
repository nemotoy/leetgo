package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func xorOperation(n int, start int) int {
	ret := 0
	for i := 0; i < n; i++ {
		ret ^= start + 2*i
	}
	return ret
}

func TestXorOperation(t *testing.T) {
	tests := []struct {
		in1 int
		in2 int
		out int
	}{
		{5, 0, 8},
		{4, 3, 8},
		{1, 7, 7},
		{10, 5, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d", tt.in1, tt.in2), func(t *testing.T) {
			got := xorOperation(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
