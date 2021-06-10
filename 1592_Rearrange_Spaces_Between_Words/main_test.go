package main

import (
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
*/
func reorderSpaces(text string) string {
	n := strings.Count(text, " ")
	if n < 1 {
		return text
	}
	tt := strings.Fields(text)
	if len(tt) == 1 {
		return tt[0] + strings.Repeat(" ", n)
	}
	plus := false
	n /= (len(tt) - 1)
	if strings.Count(text, " ")%(len(tt)-1) != 0 {
		plus = true
	}
	ret := strings.Join(tt, strings.Repeat(" ", n))
	if plus {
		ret += " "
	}
	return ret
}

func TestReorderSpaces(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"  this   is  a sentence ", "this   is   a   sentence",
		},
		{
			" practice   makes   perfect", "practice   makes   perfect ",
		},
		{
			"hello   world", "hello   world",
		},
		{
			"a", "a",
		},
		{
			"  hello", "hello  ",
		},
		{
			"a b   c d", "a b c d  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := reorderSpaces(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
