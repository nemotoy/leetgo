package main

import (
	"reflect"
	"testing"
)

/*
	## summary
	エクセルのような列番号（項目）を取得する。
	1~26: A ~ Z
	27~52: AA ~ AZ
	53~  : BA

	1桁なら1~26
	2桁で、1桁目が A=1, Z=26, 26*1or26, 2桁目は純粋に番号を加算。
	AA = 27 (1 * 26 +1)
	ZY = 701 (26 * 26 + 25 )
	3桁
	AAA = ( 26 * 26 * 1 ) + ( 26 * 1 ) + 1
	→ 桁数 - 1 が乗数
*/
func titleToNumber(s string) int {

	if len(s) == 1 {
		for _, r := range s {
			return alphabetToNumber(r)
		}
	}
	nums := make([]int, 0, len(s))
	for _, r := range s {
		num := alphabetToNumber(r)
		nums = append(nums, num)
	}
	result := 0
	len := len(nums)
	// 1の位以外は、 nの桁数-1乗して加算する。
	for i, n := range nums {
		// 1の位
		if i == len-1 {
			result += n
			break
		}
		count := len - (1 + i)
		tmp := multiplier(lettersLen, count)
		result += tmp * n
	}
	return result
}

const alphabetLetter = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lettersLen = len(alphabetLetter)

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
