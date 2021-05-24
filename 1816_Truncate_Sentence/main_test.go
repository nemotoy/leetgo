package main

import (
	"fmt"
	"strings"
	"testing"
)

func truncateSentence(s string, k int) string {
	ss := strings.Split(s, " ")[:k]
	return strings.Join(ss, " ")
}

func TestTruncateSentence(t *testing.T) {
	tests := []struct {
		in  string
		in2 int
		out string
	}{
		{
			"Hello how are you Contestant", 4, "Hello how are you",
		},
		{
			"What is the solution to this problem", 4, "What is the solution",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := truncateSentence(tt.in, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
