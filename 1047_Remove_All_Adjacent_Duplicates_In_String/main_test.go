//nolint:gocritic
package main

import (
	"fmt"
	"reflect"
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
			got := removeDuplicates(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
