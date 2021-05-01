package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func selfDividingNumbers(left int, right int) []int {
	var check = func(n int) bool {
		s := strconv.Itoa(n)
		for _, r := range s {
			i, _ := strconv.Atoi(string(r))
			if r == '0' || n%i != 0 {
				return false
			}
		}
		return true
	}

	var result []int
	for i := left; i <= right; i++ {
		if check(i) {
			result = append(result, i)
		}
	}
	return result
}

func TestSelfDividingNumbers(t *testing.T) {
	tests := []struct {
		in  int
		inN int
		out []int
	}{
		{
			1,
			22,
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.inN), func(t *testing.T) {
			got := selfDividingNumbers(tt.in, tt.inN)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
