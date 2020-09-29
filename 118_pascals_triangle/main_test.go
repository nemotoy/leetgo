package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func generate(numRows int) [][]int {
	return nil
}

func TestGenerate(t *testing.T) {
	tests := []struct {
		in  int
		out [][]int
	}{
		{
			5,
			[][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := generate(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
