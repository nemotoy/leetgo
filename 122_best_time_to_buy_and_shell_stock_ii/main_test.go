package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	in:要素が価格、インデックスが日の配列
	out: 最大利益
	基本的には安い時に買って、高い時に売るが本筋。

	- 配列が降順の場合は 0
	- 配列が昇順の場合は r-l
	- 配列の最大値と最小値を求める
		- indexが max < min の場合はNG
	- 両端から評価して、l > r ならr-1する
*/
func maxProfit(prices []int) int {

	max, min := 0, 0
	l, r := 0, len(prices)-1
	isAsc, isDesc := true, true
	for i, p := range prices {
		if i == r {
			break
		}
		if p > prices[i+1] {
			isAsc = false
		}
		if p < prices[i+1] {
			isDesc = false
		}
	}
	if isAsc {
		return prices[r] - prices[l]
	}
	if isDesc {
		return 0
	}

	for l < r {
		switch {
		case prices[l] == prices[r]:
			l, r = l+1, r-1
		case prices[l] > prices[r]:
			max = prices[l]
			r--
		case prices[l] < prices[r]:
			min = prices[r]
			l++
		}
	}

	return max - min
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxProfit(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
