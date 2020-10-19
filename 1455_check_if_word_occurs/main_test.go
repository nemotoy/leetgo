package main

import (
	"strings"
	"testing"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	// 合致しない場合は-1
	res := -1
	// sentenceを単語を要素とした、スライスにする
	words := strings.Split(sentence, " ")
	// 単語のスライスをイテレーションし、prefixに含まれるか探索する。
	for i, w := range words {
		if strings.HasPrefix(w, searchWord) {
			return i + 1
		}
	}
	return res
}

func TestIsPrefixOfWord(t *testing.T) {
	tests := []struct {
		in1 string
		in2 string
		out int
	}{
		{
			"i love eating burger", "burg", 4,
		},
		{
			"i am tired", "you", -1,
		},
		{
			"hellohello hellohellohello", "ell", -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in1+tt.in2, func(t *testing.T) {
			got := isPrefixOfWord(tt.in1, tt.in2)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
