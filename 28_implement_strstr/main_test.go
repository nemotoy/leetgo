package main

import (
	"reflect"
	"testing"
)

/*
	## summary
*/
func strStr(haystack string, needle string) int {
	return 0
}

func Test(t *testing.T) {
	tests := []struct {
		in  string
		in2 string
		out int
	}{
		{
			"hello", "ll", 2,
		},
		{
			"aaaaa", "bba", -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in+tt.in2, func(t *testing.T) {
			got := strStr(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
