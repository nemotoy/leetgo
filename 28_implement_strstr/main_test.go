package main

import (
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
*/
func strStr(haystack string, needle string) int {
	result := 0
	for _, r := range needle {
		n := strings.Count(haystack, string(r))
		if n <= 0 {
			result = -1
			break
		}
		if result < n {
			result = n
		}
	}
	return result
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
		t.Run(tt.in+"/"+tt.in2, func(t *testing.T) {
			got := strStr(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
