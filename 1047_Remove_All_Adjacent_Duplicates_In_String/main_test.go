//nolint:gocritic
package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
	2つの隣接する等しい文字を削除する。できなくなるまで繰り返す。
	- i と i+1の要素を比較する
	- 削除したら0番目から始める
	- 繰り返す※最大長が変わる
*/
func removeDuplicates(S string) string {
	i, l := 0, len(S)
	for {
		if i+1 == l {
			break
		}
		if S[i] == S[i+1] {
			if len(S) == 2 {
				S = ""
				break
			}
			S = S[:i] + S[i+2:]
			i, l = 0, len(S)
			continue
		}
		i++
	}
	return S
}

// runeの配列を初期化して、inputの文字列の各要素を評価し、同じならば削除、それ以外は追加。
func removeDuplicates2(S string) string {
	var builder []rune
	for _, char := range S {
		if notEmpty(builder) && back(builder) == char {
			builder = pop(builder)
		} else {
			builder = append(builder, char)
		}
	}
	var res strings.Builder
	for _, char := range builder {
		res.WriteRune(char)
	}
	return res.String()
}

func back(arr []rune) rune {
	return arr[len(arr)-1]
}

func notEmpty(arr []rune) bool {
	return len(arr) > 0
}

func pop(arr []rune) []rune {
	if notEmpty(arr) {
		return arr[:len(arr)-1]
	}
	return arr
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"abbaca", "ca",
		},
		{
			"aaaaaaaa", "",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := removeDuplicates2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
