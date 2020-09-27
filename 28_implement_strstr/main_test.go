package main

import (
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
	needleが存在するindexを返す。
*/
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
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
		{
			"a", "a", 0,
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
