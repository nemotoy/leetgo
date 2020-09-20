package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func climbStairs(n int) int {
	return 0
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			2, 2,
		},
		{
			3, 3,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := climbStairs(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
