package main

import (
	"fmt"
	"strings"
	"testing"
)

/*
	## summary
	スペース以外の文字の連続したシーケンスの個数
*/
func countSegments(s string) int {
	ret := 0
	for _, s := range strings.Split(s, " ") {
		if len(s) > 0 {
			ret++
		}
	}
	return ret
}

func TestCountSegments(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"Hello, my name is John", 5,
		},
		{
			"Hello", 1,
		},
		{
			"", 0,
		},
		{
			"                ", 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countSegments(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
