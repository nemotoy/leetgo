package main

import (
	"reflect"
	"testing"
)

/*
	## summary
*/
func isValid(s string) bool {
	return false
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"()", true,
		},
		{
			"()[]{}", true,
		},
		{
			"(]", false,
		},
		{
			"(]", false,
		},
		{
			"{[]}", true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := isValid(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
