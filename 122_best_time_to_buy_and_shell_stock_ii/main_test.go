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
	- 複数回購入・売却する場合
*/
func maxProfit(prices []int) int {

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

	prof := 0
	for i, p := range prices {
		if i == r-1 {
			break
		}

		// 山の登斜を見つける
		if p < prices[i+1] {
			next := 2
			for {
				fmt.Printf("val: %d, inc: %d, len: %d\n", p, i, next)
				switch {
				case r-i < next:
					fmt.Printf("#1 Profit: %d, Add: %d\n", prof, prices[i+next-1]-prices[i+1])
					// prof += prices[i+next-1] - prices[i+1]
				case prices[i+next-1] < prices[i+next]: // 登斜の場合は継続
					next++
					continue
				case prices[i+next-1] > prices[i+next]: // 降斜の場合は利益決定
					fmt.Printf("#3 Profit: %d, Add: %d\n", prof, prices[i+next-1]-prices[i+1])
					prof += prices[i+next-1] - prices[i+1]
				}
				break
			}
		}
	}
	return prof
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
		// {
		// 	[]int{1, 3, 2, 2, 4, 1, 5},
		// 	8,
		// },
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
