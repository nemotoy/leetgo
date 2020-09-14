package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
	## summary
	正の整数の各要素の二乗の合計が1になる場合、trueを返す。

	> it loops endlessly in a cycle which does not include 1.
	無限ループをどのように評価するのか。
*/
func isHappy(n int) bool {
	// 負の整数はfalseを返す
	if n < 1 {
		return false
	}
	// 1桁ずつ扱うために、文字列に型変換する。
	s := strconv.Itoa(n)
	// i < 10 の 10 は無限ループを回避するための適当な値。
	for i := 0; i < 10; i++ {
		amount := 0
		// 各桁を取得し、int型に変換し、二乗した値をamountに代入する。
		for _, r := range s {
			nn, _ := strconv.Atoi(string(r))
			amount += nn * nn
		}
		if amount == 1 {
			return true
		}
		// 次回のiterationでは今回算出した合計値を利用する
		s = strconv.Itoa(amount)
	}
	return false
}

func TestIsHappy(t *testing.T) {
	tests := []struct {
		in  int
		out bool
	}{
		{
			19, true,
		},
		{
			7, true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := isHappy(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
