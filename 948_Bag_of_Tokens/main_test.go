package main

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary

	1. if power => tokens[i]
	power -= tokens[i], score += 1
	2. if score >= 1
	power += tokens[i], score -= 1

	各トークンは最大で1回、任意の順序で再生できる。 すべてのトークンをプレイする必要はない。
	最大スコアを返す。
	表は小さいトークン、裏は大きいトークンが定石。

	- 昇順ソート
	- プレイ
	  - 表向きでプレイ継続
	   - 裏向きでプレイ

*/
// nolint
func bagOfTokensScore(tokens []int, P int) int {
	// sort asc
	sort.Ints(tokens)
	indx, length := 0, len(tokens)-1
	score, ans := 0, 0
	for indx <= length && (P >= tokens[indx] || score > 0) {
		// head: 条件に合致し続けるまで処理
		for indx <= length && P >= tokens[indx] {
			P -= tokens[indx]
			score++
			indx++
		}
		// headの処理結果と前回の処理結果の大きい値を暫定値とする。
		ans = max(ans, score)
		// tail:
		if indx <= length && score > 0 {
			P += tokens[length]
			score--
			length--
		}
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestBagOfTokensScore(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{
			[]int{100}, 50, 0,
		},
		{
			[]int{100, 200}, 150, 1,
		},
		{
			[]int{100, 200, 300, 400}, 200, 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := bagOfTokensScore(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
