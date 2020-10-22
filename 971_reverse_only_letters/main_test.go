package main

import (
	"testing"
	"unicode"
)

/*
	## summary
*/
func reverseOnlyLetters(s string) string {
	res := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !unicode.IsLetter(rune(res[i])) {
			i++
		}
		for i < j && !unicode.IsLetter(rune(res[j])) {
			j--
		}
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return string(res)
}

func TestReverseOnlyLetters(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"ab-cd", "dc-ba",
		},
		{
			"a-bC-dEf-ghIj", "j-Ih-gfE-dCba",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := reverseOnlyLetters(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
