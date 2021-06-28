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
	for _, v := range image {
		// flip
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
		// invert
		for k := 0; k < len(v); k++ {
			v[k] = invert(v[k])
		}
	}
	return image
}

func invert(n int) int {
	if n == 0 {
		return 1
	}
	return 0
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
				{0, 1, 1, 0},
				{0, 0, 0, 1},
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
