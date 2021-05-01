package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func flipAndInvertImage(image [][]int) [][]int {
	return [][]int{}
}

func TestFlipAndInvertImage(t *testing.T) {
	tests := []struct {
		in  [][]int
		out [][]int
	}{
		{
			[][]int{
				{1, 1, 0},
				{1, 0, 1},
				{0, 0, 0},
			},
			[][]int{
				{1, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
		},
		{
			[][]int{
				{1, 1, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
			[][]int{
				{1, 1, 0, 0},
				{0, 1, 1, 1},
				{0, 1, 1, 1},
				{1, 0, 1, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := flipAndInvertImage(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
