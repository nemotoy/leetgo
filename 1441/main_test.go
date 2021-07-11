package main

import (
	"reflect"
	"testing"
)

/*
  ## summary
*/
func buildArray(target []int, n int) []string {
	ret := make([]string, 0, n*2)
	for i, j := 1, 0; i <= n && j < len(target); i++ {
		ret = append(ret, "Push")
		if target[j] == i {
			j++
		} else {
			ret = append(ret, "Pop")
		}
	}
	return ret
}

func TestBuildArray(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out []string
	}{
		{
			[]int{1, 3},
			3,
			[]string{"Push", "Push", "Pop", "Push"},
		},
		{
			[]int{1, 2, 3},
			3,
			[]string{"Push", "Push", "Push"},
		},
		{
			[]int{1, 2},
			4,
			[]string{"Push", "Push"},
		},
	}
	for _, tt := range tests {
		got := buildArray(tt.in, tt.in2)
		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("got: %v, want: %v", got, tt.out)
		}
	}
}
