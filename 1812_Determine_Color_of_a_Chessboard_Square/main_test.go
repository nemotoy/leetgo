package main

import (
	"testing"
)

/*
  ## summary
  - x: a ~ h(black ~ white)
  - y: 1 ~ 8(black ~ white)

  Return true if the square is white, and false if the square is black.

  (奇数、奇数), （偶数、偶数）の場合、black
*/
func squareIsWhite(coordinates string) bool {
	x, y := coordinates[0], coordinates[1]
	return (x%2 != 0 && y%2 == 0) || (x%2 == 0 && y%2 != 0)
}

func TestSquareIsWhite(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"a1", false,
		},
		{
			"h3", true,
		},
		{
			"c7", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := squareIsWhite(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
