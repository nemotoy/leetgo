package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
	スペース区切りの文字列の最後の文字数
*/
func lengthOfLastWord(s string) int {
	ss := strings.Split(strings.TrimSpace(s), " ")
	return len(ss[len(ss)-1])
}

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"Hello World", 5,
		},
		{
			" ", 0,
		},
		{
			"a ", 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := lengthOfLastWord(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
