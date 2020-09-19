package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	k >= 0, 配列の右からk分左に移動する。
	配列のインデックスをずらすことができれば、メモリアロケートしなくて済む。
	例）
	array = [1,2,3,4], k = 2 の期待値は [3,4,1,2]なので、配列の前からk-1番目まではインデックスが +k され、配列後からk番目まではインデックスが -k される。
	- l-k > k 場合は通らない。
	1. l < kまで、最後尾の要素を先頭のインデックスに変更する。
*/
func rotate(nums []int, k int) {
	l := len(nums)
	if l == 0 {
		return
	}
	for i := 0; i < k; i++ {
		for j := l - 1; j > 0; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j] // 最後尾の要素を先頭のインデックスに変更する。
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		in1 []int
		in2 int
		out []int
	}{
		{
			[]int{1, 2, 3, 4},
			2,
			[]int{3, 4, 1, 2},
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("nums: %v, k: %d", tt.in1, tt.in2), func(t *testing.T) {
			rotate(tt.in1, tt.in2)
			got := tt.in1
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
