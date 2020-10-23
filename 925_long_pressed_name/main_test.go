package main

import (
	"reflect"
	"testing"
)

/*
	## summary
*/
func isLongPressedName(name string, typed string) bool {
	// j represents an index of the typed.
	var j, cnt int
	for i := 0; i < len(name); i++ {
		if j == len(typed) {
			return false
		}
		if name[i] == typed[j] {
			cnt++
		} else {
			i--
		}
		j++
	}
	return cnt == len(name)
}

func TestIsLongPressedName(t *testing.T) {
	tests := []struct {
		in  string
		in2 string
		out bool
	}{
		{
			"alex", "aaleex", true,
		},
		{
			"saeed", "ssaaedd", false,
		},
		{
			"leelee", "lleeelee", true,
		},
		{
			"laiden", "laiden", true,
		},
		{
			"kikcxmvzi", "kiikcxxmmvvzz", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in+tt.in2, func(t *testing.T) {
			got := isLongPressedName(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
