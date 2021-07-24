package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	隣接せずにn個の'1'を配置することが可能か
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			count++
		}
	}
	return count >= n
}

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out bool
	}{
		{
			[]int{1, 0, 0, 0, 1}, 1, true,
		},
		{
			[]int{1, 0, 0, 0, 1}, 2, false,
		},
		{
			[]int{1, 0, 0, 0, 0, 0, 1}, 2, true,
		},
		{
			[]int{1, 0, 0, 0, 0, 1}, 2, false,
		},
		{
			[]int{1, 0, 1, 0, 1, 0, 1}, 1, false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := canPlaceFlowers(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
