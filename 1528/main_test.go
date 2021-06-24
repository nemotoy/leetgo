package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func restoreString(s string, indices []int) string {
	ret := make([]byte, len(s))
	for i, n := range indices {
		ret[n] = s[i]
	}
	return string(ret)
}

func TestRestoreString(t *testing.T) {
	tests := []struct {
		in  string
		in2 []int
		out string
	}{
		{"codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
		{"abc", []int{0, 1, 2}, "abc"},
		{"aiohn", []int{3, 1, 4, 2, 0}, "nihao"},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := restoreString(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
