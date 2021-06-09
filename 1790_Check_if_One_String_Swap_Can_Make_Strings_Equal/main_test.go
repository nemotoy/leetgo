package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	最大1回のスワップで同値になるか
*/
func areAlmostEqual(s1 string, s2 string) bool {
	chars := make([]int, 0, 2)
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			chars = append(chars, i)
		}
		if len(chars) >= 3 {
			return false
		}
	}
	return len(chars) == 0 || len(chars) == 2 && s1[chars[0]] == s2[chars[1]] && s1[chars[1]] == s2[chars[0]]
}

func TestAreAlmostEqual(t *testing.T) {
	tests := []struct {
		in1 string
		in2 string
		out bool
	}{
		{
			"bank", "kanb", true,
		},
		{
			"attack", "defend", false,
		},
		{
			"kelb", "kelb", true,
		},
		{
			"aa", "ac", false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %s", tt.in1, tt.in2), func(t *testing.T) {
			got := areAlmostEqual(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
