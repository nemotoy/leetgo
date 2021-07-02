package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
// nolint:gocritic
// map
func repeatedNTimes(A []int) int {
	m := make(map[int]int, len(A))
	for _, a := range A {
		m[a] += 1
	}
	for k, v := range m {
		if v > 1 {
			return k
		}
	}
	return 0
}

// nolint:gocritic
// 1つの要素がrepeatedということは、他の要素はnon repeatedということ。
func repeatedNTimes2(A []int) int {
	for i := 2; i < len(A); i++ {
		if A[i] == A[i-1] || A[i] == A[i-2] {
			return A[i]
		}
	}
	return A[0]
}

func TestRepeatedNTimes(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{1, 2, 3, 3}, 3},
		{[]int{2, 1, 2, 5, 3, 2}, 2},
		{[]int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := repeatedNTimes2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
