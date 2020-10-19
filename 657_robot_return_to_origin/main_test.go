package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func judgeCircle(moves string) bool {
	var x, y int
	for _, c := range moves {
		if c == 'L' {
			x--
		}
		if c == 'R' {
			x++

		}
		if c == 'U' {
			y++
		}
		if c == 'D' {
			y--
		}
	}
	return x == 0 && y == 0
}

func TestJudgeCircle(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"UD", true,
		},
		{
			"LL", false,
		},
		{
			"RRDD", false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := judgeCircle(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
