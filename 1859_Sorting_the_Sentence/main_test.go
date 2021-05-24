package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func sortSentence(s string) string {
	// " "毎に要素を分割し配列にする
	ss := strings.Split(s, " ")
	// 結果の初期化
	ret := make([]string, len(ss))
	for _, s := range ss {
		// 接尾の数値を取得する
		n, _ := strconv.Atoi(string(s[len(s)-1]))
		// index（数値-1）に数値以外の文字列を追加する
		ret[n-1] = s[:len(s)-1]
	}
	return strings.Join(ret, " ")
}

func TestSortSentence(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"is2 sentence4 This1 a3", "This is a sentence",
		},
		{
			"Myself2 Me1 I4 and3", "Me Myself and I",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sortSentence(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
