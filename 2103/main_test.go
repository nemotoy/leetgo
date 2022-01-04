package main

import (
	"testing"
)

/*
	## summary
	RGBが全て含まれるrod数
*/
func countPoints(rings string) int {
	// color:rod
	cnt := [10][3]int{{}}
	for i := 0; i < len(rings); i += 2 {
		color, rod := convertColorToNum(rings[i]), rings[i+1]-'0'
		cnt[rod][color] = 1
	}
	ret := 0
	for i := 0; i < 10; i++ {
		if cnt[i][0]+cnt[i][1]+cnt[i][2] == 3 {
			ret++
		}
	}
	return ret
}

func convertColorToNum(s byte) int {
	color := 0
	switch s {
	case 'R':
		color = 0
	case 'G':
		color = 1
	case 'B':
		color = 2
	}
	return color
}

func TestCountPoints(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"B0B6G0R6R0R6G9", 1,
		},
		{
			"B0R0G0R9R0B0G0", 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := countPoints(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
