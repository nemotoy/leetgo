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

	- 購入・売却・利益・サイズを定義する
	- 始点がサイズを超過するまでイテレートする
		- 山の登斜を見つけるまで購入始点を進める
		- 同始点まで売却始点を進める
		- 山の降斜を見つけるまで売却始点を進める
		- 利益を算出する
		- 購入始点を売却始点の次点に変更する。
*/
func maxProfit(prices []int) int {
	buy := 0
	sell := 0
	profit := 0
	n := len(prices)
	for buy < n && sell < n {
		for buy+1 < n && prices[buy+1] < prices[buy] {
			buy++
		}
		sell = buy
		for sell+1 < n && prices[sell+1] > prices[sell] {
			sell++
		}

		profit += prices[sell] - prices[buy]
		buy = sell + 1
	}
	return profit
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
		{
			[]int{6, 1, 3, 2, 4, 7}, // 1. buy day2(1) sell day3(3) profit = 2; 2. buy day4(2) sell day6(7) profit = 5; amount = 7
			7,
		},
		{
			[]int{1, 4, 2},
			3,
		},
		{
			[]int{2, 1, 4},
			3,
		},
		{
			[]int{5, 2, 3, 2, 6, 6, 2, 9, 1, 0, 7, 4, 5, 0},
			20,
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
