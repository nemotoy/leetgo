package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	- 基本的に大きい数字が左から並ぶ
		- 4はIIIIではなく、IVと書く
*/
func romanToInt(s string) int {
	result := 0
	for _, r := range s {
		v, ok := romanNumerals[string(r)]
		if ok {
			result += v
		}
	}
	return result
}

var romanNumerals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"III",
			3,
		},
		{
			"IV",
			4,
		},
		{
			"IX",
			9,
		},
		{
			"LVIII",
			58,
		},
		{
			"MCMXCIV",
			1994,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := romanToInt(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
