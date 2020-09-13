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

	I can be placed before V (5) and X (10) to make 4 and 9.
	X can be placed before L (50) and C (100) to make 40 and 90.
	C can be placed before D (500) and M (1000) to make 400 and 900.

	基底indexの次の要素を取り出し、上の条件に合致したら、計算し、次次点のincrementから再開する。
*/
func romanToInt(s string) int {
	result := 0
	size := len(s)
	for i := 0; i < size; i++ {
		ele := string(s[i])
		v, ok := romanNumerals[string(ele)]
		if ok {
			if i+1 < size {
				switch ele {
				case "I":
					if string(s[i+1]) == "V" {
						v = 4
						i++
						fmt.Println("#I", v, s)
					} else if string(s[i+1]) == "X" {
						v = 9
						i++
					}
				case "X":
					if string(s[i+1]) == "L" {
						v = 40
						i++
					} else if string(s[i+1]) == "C" {
						v = 90
						i++
					}
				case "C":
					if string(s[i+1]) == "D" {
						v = 400
						i++
					} else if string(s[i+1]) == "M" {
						v = 900
						i++
					}
				}
			}
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

func TestRomanToInt(t *testing.T) {
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
