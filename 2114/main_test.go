package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	文に含まれるワードの最大数
*/
func mostWordsFound(sentences []string) int {
	ret := 0
	for _, sentence := range sentences {
		ss := strings.Split(sentence, " ")
		l := len(ss)
		if ret < l {
			ret = l
		}
	}
	return ret
}

func TestMostWordsFound(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{[]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}, 6},
		{[]string{"please wait", "continue to fight", "continue to win"}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := mostWordsFound(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
