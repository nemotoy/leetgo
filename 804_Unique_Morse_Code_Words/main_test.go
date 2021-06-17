package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func uniqueMorseRepresentations(words []string) int {
	var morse = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	ret := map[string]struct{}{}
	for _, word := range words {
		code := ""
		for _, w := range word {
			code += morse[w-'a']
		}
		ret[code] = struct{}{}
	}
	return len(ret)
}

func TestUniqueMorseRepresentations(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{
			[]string{"gin", "zen", "gig", "msg"}, 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := uniqueMorseRepresentations(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
