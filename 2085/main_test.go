package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	両文字列の配列で一度ずつ出現する文字列の総数
*/
func countWords(words1 []string, words2 []string) int {
	m1 := make(map[string]int, len(words1))
	m2 := make(map[string]int, len(words2))
	for _, word1 := range words1 {
		m1[word1]++
	}
	for _, word2 := range words2 {
		m2[word2]++
	}

	from := words1
	if len(words1) > len(words2) {
		from = words2
	}

	ret := 0
	for _, w := range from {
		if m1[w] == 1 && m2[w] == 1 {
			ret++
		}
	}
	return ret
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		in  []string
		in2 []string
		out int
	}{
		{
			[]string{"leetcode", "is", "amazing", "as", "is"},
			[]string{"amazing", "leetcode", "is"},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.in, tt.in2), func(t *testing.T) {
			got := countWords(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
