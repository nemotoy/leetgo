package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func countPrimes(n int) int {
	return 0
}

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			10, 4,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := countPrimes(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
