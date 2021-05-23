package main

import (
	"strings"
	"testing"
)

func checkIfPangram(sentence string) bool {
	// shorter than length of English alphabet letters
	if len(sentence) < 26 {
		return false
	}
	alps := "abcdefghijklmnopqrstuvwxyz"
	for _, s := range alps {
		if !strings.ContainsRune(sentence, s) {
			return false
		}
	}
	return true
}

func TestCheckIfPangram(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"thequickbrownfoxjumpsoverthelazydog", true,
		},
		{
			"leetcode", false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := checkIfPangram(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
