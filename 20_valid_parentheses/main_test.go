package main

import (
	"reflect"
	"testing"
)

/*
	## summary
*/
func isValid(s string) bool {
	stack := []string{}
	maps := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	for _, r := range s {
		if v, ok := maps[string(r)]; ok {
			top := "#"
			if len(stack) != 0 {
				tail := len(stack) - 1
				top = stack[tail]
				stack = stack[:tail]
			}
			if v != top {
				return false
			}
		} else {
			stack = append(stack, string(r))
		}
	}
	return len(stack) == 0
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
