package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func findGCD(nums []int) int {
	sort.Ints(nums)
	max, min := nums[len(nums)-1], nums[0]
	return gcd(max, min)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func TestFindGCD(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{2, 5, 6, 9, 10}, 2},
		{[]int{7, 5, 6, 8, 3}, 1},
		{[]int{3, 3}, 3},
		{[]int{6, 7, 9}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findGCD(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
