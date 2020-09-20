package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	階段の上階に登る。一度に1か2階段登れる。取り得る選択肢を返す。
	- 全て1の場合
	- 2が含まれる場合
		- 個数は要素の長さ
	- 全て2の場合
*/
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	r := 0
	r += 1 // 全て1の場合
	// 全て2の場合
	if n%2 == 0 {
		r += 1
	}
	l := n - 2 + 1
	// 2が含まれる場合
	for n-2 > 0 {
		r += l
		n -= 2
		l--
	}
	return r
}

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			2, 2,
		},
		{
			3, 3,
		},
		{
			4, 5,
		},
		{
			5, 8,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := climbStairs(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
