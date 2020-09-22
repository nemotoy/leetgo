package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	isBadVersion() APIの呼び出し回数を出来るだけ少なくする。
	一度trueになったら、それ以降のnはいずれもtrueを返す。という性質を利用するはず。(false -> true -> true) を見つける。

	1. 1からnまでイテレーションし、それぞれAPIコールする
	2. l/rの中間点を探索する
		- l/rの中間点が
			- trueなら、左辺を探索
			- falseなら、右辺を探索
*/
func firstBadVersion(n int) int {
	// nは1からnまでなので、基底を1で開始し、nまで。
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return 0
}

func firstBadVersion2(n int) int {
	left, right := 1, n
	for left < right {
		mid := (left + right) / 2

		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func isBadVersion(version int) bool {
	return version > 5 // 適当なロジック。与件ではない。
}

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := firstBadVersion(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
