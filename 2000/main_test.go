package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	chがある場合、0からchを含んだindexまで反転。なければしない。
*/
func reversePrefix(word string, ch byte) string {
	for i := 0; i < len(word); i++ {
		if word[i] == ch {
			return reverse(word[0:i+1]) + word[i+1:]
		}
	}
	return word
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TestReversePrefix(t *testing.T) {
	tests := []struct {
		in  string
		in2 byte
		out string
	}{
		{
			"abcdefd", 'd', "dcbaefd",
		},
		{
			"xyxzxe", 'z', "zxyxxe",
		},
		{
			"abcd", 'z', "abcd",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := reversePrefix(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
