package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	ret, n := 0, len(seats)
	for i := 0; i < n; i++ {
		ret += abs(seats[i] - students[i])
	}
	return ret
}

func abs(x int) int {
	if x < 1 {
		return -x
	}
	return x
}

func TestMinMovesToSeat(t *testing.T) {
	tests := []struct {
		in  []int
		in2 []int
		out int
	}{
		{[]int{3, 1, 5}, []int{2, 7, 4}, 4},
		{[]int{4, 1, 5, 9}, []int{1, 3, 2, 6}, 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.in2), func(t *testing.T) {
			got := minMovesToSeat(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
