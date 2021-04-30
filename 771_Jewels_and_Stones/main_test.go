package main

import (
	"fmt"
	"testing"
)

/*
	## summary
*/
func numJewelsInStones(jewels string, stones string) int {
	var result int
	for i := 0; i < len(jewels); i++ {
		for ii := 0; ii < len(stones); ii++ {
			if stones[ii] == jewels[i] {
				result++
			}
		}
	}
	return result
}

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		in1 string
		in2 string
		out int
	}{
		{
			"aA", "aAAbbbb", 3,
		},
		{
			"z", "ZZ", 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %s", tt.in1, tt.in2), func(t *testing.T) {
			got := numJewelsInStones(tt.in1, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
