package main

import (
	"testing"
)

/*
	## summary
*/
func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	occurrences := m[rune(s[0])]
	for _, c := range m {
		if c != occurrences {
			return false
		}
	}
	return true
}

func TestAreOccurrencesEqual(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{"abacbc", true},
		{"aaabb", false},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := areOccurrencesEqual(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
