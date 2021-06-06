package main

import (
	"fmt"
	"testing"
)

/*
	## summary
	in: time(hh:mm) some of digits in the string are hidden(represented by ?)
	out: time
*/
// nolint:gocritic
func maximumTime(time string) string {
	ret := []byte(time)
	if time[0] == '?' && time[1] == '?' {
		ret[0] = '2'
		ret[1] = '3'
	} else if time[0] == '?' {
		if time[1] < '4' {
			ret[0] = '2'
		} else {
			ret[0] = '1'
		}
	} else if time[1] == '?' {
		if time[0] == '2' {
			ret[1] = '3'
		} else {
			ret[1] = '9'
		}
	}
	if time[3] == '?' && time[4] == '?' {
		ret[3] = '5'
		ret[4] = '9'
	} else if time[3] == '?' {
		if time[4] != '0' {
			ret[3] = '5'
		} else {
			ret[3] = '5'
		}
	} else if time[4] == '?' {
		if time[3] == '6' {
			ret[4] = '0'
		} else {
			ret[4] = '9'
		}
	}
	return string(ret)
}

func TestMaximumTime(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"2?:?0", "23:50",
		},
		{
			"0?:3?", "09:39",
		},
		{
			"1?:22", "19:22",
		},
		{
			"?0:15", "20:15",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maximumTime(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
