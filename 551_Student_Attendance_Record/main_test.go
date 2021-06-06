package main

import (
	"testing"
)

/*
	## summary
	- Aが2未満
	- Lが3以上連続しない
*/
func checkRecord(s string) bool {
	a, l := 0, 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			a++
			l = 0
		case 'L':
			l++
		default:
			l = 0
		}
		if a >= 2 || l > 2 {
			return false
		}
	}
	return true
}

func TestCheckRecord(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"PPALLP", true,
		},
		{
			"PPALLL", false,
		},
		{
			"LALL", true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := checkRecord(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
