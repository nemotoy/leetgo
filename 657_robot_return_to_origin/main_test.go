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
	var ud, lr int
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'U':
			ud += 1
		case 'D':
			ud -= 1
		case 'L':
			lr += 1
		case 'R':
			lr -= 1
		}
	}
	return ud == 0 && lr == 0
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
