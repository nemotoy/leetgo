package main

import (
	"fmt"
	"strconv"
	"testing"
)

/*
  ## summary
  in: lowercase letters and/or digits
  out: second largest numerical digit that appears in s or -1 if it does not exist
*/
func secondHighest(s string) int {
	first, second := -1, -1
	for i := 0; i < len(s); i++ {
		n, err := strconv.Atoi(string(s[i]))
		if err == nil {
			// s is a numeric.
			if n > first {
				second, first = first, n
			} else if n < first && n > second {
				second = n
			}
		}
	}
	return second
}

func secondHighest2(s string) int {
	first, second := -1, -1
	for _, r := range s {
		if r < '0' || r > '9' {
			continue
		}
		n := int(r - '0')
		if n > first {
			second, first = first, n
		} else if n < first && n > second {
			second = n
		}
	}
	return second
}

func isNumeric(r rune) bool {
	return r < '0' || r > '9'
}

func TestSecondHighest(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"dfa12321afd", 2,
		},
		{
			"abc1111", -1,
		},
		{
			"ck077", 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := secondHighest2(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
