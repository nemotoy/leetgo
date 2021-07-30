package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	out: トップに到達する最低コスト
	cost[i] がコスト。コストを支払うと1か2つ加算できる
	cost[0] か cost[1] から開始できる
*/
func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		in  []int
		out int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minCostClimbingStairs(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
