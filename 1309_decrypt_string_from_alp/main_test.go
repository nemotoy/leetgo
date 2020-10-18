package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func freqAlphabets(s string) string {
	r := ""
	for i := 0; i < len(s); {
		// 現在のindexから#が出現するのは2要素先。(範囲は 10#〜26#)
		// 探索対象が3要素以下だった場合panicするので、まず長さを評価する。
		// 10#〜26#の場合は以下の処理を行う。
		if i+2 < len(s) && string(s[i+2]) == "#" {
			// 1桁目。10を乗算しているのは、10の位として扱うため。
			// 0のruneを引いているのは、コードポイントの計算を1~9にならすため。
			l1 := (s[i] - '0') * 10
			// 2桁目
			l2 := (s[i+1] - '0')
			ss := l1 + l2
			// 結果の文字列型に結合
			r += string(ss + 'a' - 1)
			// #まで評価済みのため、indexを進める。
			i += 3
		} else {
			r += string(s[i] - '0' - 1 + 'a')
			i++
		}
	}
	return r
}

func TestFreqAlphabets(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"10#11#12", "jkab",
		},
		{
			"1326#", "acz",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := freqAlphabets(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
