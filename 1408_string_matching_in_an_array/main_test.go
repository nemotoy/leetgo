package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
*/
func stringMatching(words []string) []string {
	hashTable := map[string]bool{}
	res := []string{}
	for _, v := range words {
		for _, v2 := range words {
			if v == v2 || len(v2) > len(v) {
				continue
			}
			if _, ok := hashTable[v2]; !ok && strings.Contains(v, v2) {
				hashTable[v2] = true
				res = append(res, v2)
			}
		}
	}
	return res
}

func TestStringMatching(t *testing.T) {
	tests := []struct {
		in  []string
		out []string
	}{
		{
			[]string{"mass", "as", "hero", "superhero"}, []string{"as", "hero"},
		},
		{
			[]string{"leetcode", "et", "code"}, []string{"et", "code"},
		},
		{
			[]string{"leetcoder", "leetcode", "od", "hamlet", "am"}, []string{"leetcode", "od", "am"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := stringMatching(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
