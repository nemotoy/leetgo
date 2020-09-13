package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func romanToInt(s string) int {
	return 0
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
