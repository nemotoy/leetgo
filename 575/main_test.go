package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	return the maximum number of different types of candies she can eat if she only eats n / 2 of them.
*/
func distributeCandies(candyType []int) int {
	canEatNumber := len(candyType) / 2
	uniqueCandies := make(map[int]int)
	for _, candie := range candyType {
		uniqueCandies[candie] = 1
	}
	return min(len(uniqueCandies), canEatNumber)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestDistributeCandies(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{1, 1, 2, 2, 3, 3}, 3,
		},
		{
			[]int{1, 1, 2, 3}, 2,
		},
		{
			[]int{6, 6, 6, 6}, 1,
		},
		{
			[]int{100, 1, 1, 1}, 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := distributeCandies(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
