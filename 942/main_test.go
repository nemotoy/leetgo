package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  ## summary
*/
func diStringMatch(s string) []int {
	n := len(s)
	ret := make([]int, n+1)
	left, right := 0, n
	for i := 0; i < n; i++ {
		if s[i] == 'I' {
			ret[i] = left
			left++
		} else {
			ret[i] = right
			right--
		}
	}
	ret[n] = left
	return ret
}

func TestDiStringMatch(t *testing.T) {
	tests := []struct {
		in  string
		out []int
	}{
		{"IDID", []int{0, 4, 1, 3, 2}},
		{"III", []int{0, 1, 2, 3}},
		{"DDI", []int{3, 2, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := diStringMatch(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
