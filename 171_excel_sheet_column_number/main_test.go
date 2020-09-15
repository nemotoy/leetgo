package main

import (
	"reflect"
	"testing"
)

/*
	## summary
	エクセルのような列番号（項目）を取得する。

	アルファベット数字をnとした場合
	- 1桁の場合: n
	- 2桁以上の場合: 26を桁数-1 回が乗算し、さらに n を乗算し、1の位の数字を加算する
	AA = 27 (26 * 1 * 1) + 1
	ZY = 701 (26 * 26 * 1) + 25
	AAA = 703 (26 * 26 * 1 * 1) + ( 26 * 1 * 1) + 1
*/
func titleToNumber(s string) int {

	len := len(s)
	if len == 1 {
		for _, r := range s {
			return alphabetToNumber(r)
		}
	}
	result := 0
	for i, r := range s {
		n := alphabetToNumber(r)
		if i == len-1 {
			result += n
			break
		}
		result += multiplier(lettersLen, len-(1+i)) * n
	}

	return result
}

const alphabetLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettersLen = len(alphabetLetter)

// nの乗数（count）を返す
func multiplier(n, count int) int {
	result := 1
	for count > 0 {
		result *= n
		count--
	}
	return result
}

// アルファベットを数値に変換する（1~26）
func alphabetToNumber(r rune) int {
	for i, a := range alphabetLetter {
		if r == a {
			return i + 1
		}
	}
	return 0
}

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"A", 1,
		},
		{
			"AA", 27,
		},
		{
			"AB", 28,
		},
		{
			"ZY", 701,
		},
		{
			"AAA", 703,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := titleToNumber(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
