package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func reverse(x int) int {
	return 0
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
