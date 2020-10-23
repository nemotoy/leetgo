package main

import (
	"fmt"
	"strconv"
	"testing"
)

/*
	## summary
	百位に "." を付与した文字列を返す。
	- 1位を数え "." を付与していく。
	- nの長さを基底値とし、3（3桁）を decrementしていく。
*/
func thousandSeparator(n int) string {
	// n is not thousand, return a string of n.
	if n < 1000 {
		return strconv.Itoa(n)
	}
	// delimiter letter
	dell := 3
	s := strconv.Itoa(n)
	l := len(s) - dell
	for l > 0 {
		s = s[:l] + "." + s[l:]
		l -= dell
	}
	return s
}

func TestThousandSeparator(t *testing.T) {
	tests := []struct {
		in  int
		out string
	}{
		{
			987, "987",
		},
		{
			1234, "1.234",
		},
		{
			123456789, "123.456.789",
		},
		{
			0, "0",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := thousandSeparator(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
