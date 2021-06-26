package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	return the number of rectangles that can make a square with a side length of max length.
*/
func countGoodRectangles(rectangles [][]int) int {
	squareLengths := make([]int, len(rectangles))
	// get minimum values each rectangles
	for i, regrectangle := range rectangles {
		h, w := regrectangle[0], regrectangle[1]
		squareLengths[i] = min(h, w)
	}
	ret, max := 0, 0
	// get the number of max length
	for _, length := range squareLengths {
		if max < length {
			max = length
			// because of replacing with length, resets the ret,
			ret = 0
			ret++
		} else if max == length {
			ret++
		}
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestCountGoodRectangles(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}}, 3,
		},
		{
			[][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}}, 3,
		},
		{
			[][]int{{5, 8}, {3, 9}, {3, 12}}, 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countGoodRectangles(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
