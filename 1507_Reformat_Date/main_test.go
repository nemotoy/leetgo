package main

import (
	"fmt"
	"strings"
	"testing"
)

func reformatDate(date string) string {
	dd := strings.Split(date, " ")
	// reverse a given array
	for i, j := 0, len(dd)-1; i < j; i, j = i+1, j-1 {
		dd[i], dd[j] = dd[j], dd[i]
	}
	month := map[string]string{"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06", "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12"}
	dd[1] = month[dd[1]]
	day := dd[2]
	day = day[:len(day)-2]
	if len(day) == 1 {
		day = "0" + day
	}
	dd[2] = day
	return strings.Join(dd, "-")
}

func TestReformatDate(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"20th Oct 2052", "2052-10-20",
		},
		{
			"6th Jun 1933", "1933-06-06",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := reformatDate(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
