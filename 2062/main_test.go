package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	全ての連続した母音で構成される文字シーケンス数

	brute force
*/
func countVowelSubstrings(word string) int {
	l := len(word)
	if l < 5 {
		return 0
	}
	ret := 0
	for i := 0; i < l; i++ {
		m := make(map[byte]int, 5)
		for j := i; j < l && isVowel(word[j]); j++ {
			m[word[j]]++
			if len(m) == 5 {
				ret++
			}
		}
	}
	return ret
}

func isVowel(s byte) bool {
	return s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u'
}

func TestCountVowelSubstrings(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"aeiouu",
			2,
		},
		{
			"unicornarihan",
			0,
		},
		{
			"cuaieuouac",
			7,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := countVowelSubstrings(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
