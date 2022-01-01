package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	arrに一回だけ存在する文字列で、k番目の個別の文字列を返す
*/
func kthDistinct(arr []string, k int) string {
	m := make(map[string]int, len(arr))
	for _, a := range arr {
		m[a]++
	}
	for _, a := range arr {
		if m[a] == 1 {
			k--
			if k == 0 {
				return a
			}
		}
	}
	return ""
}

func TestKthDistinct(t *testing.T) {
	tests := []struct {
		in  []string
		in2 int
		out string
	}{
		{[]string{"d", "b", "c", "b", "c", "a"}, 2, "a"},
		{[]string{"aaa", "aa", "a"}, 1, "aaa"},
		{[]string{"a", "b", "a"}, 3, ""},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v,%v", tt.in, tt.in2), func(t *testing.T) {
			got := kthDistinct(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
