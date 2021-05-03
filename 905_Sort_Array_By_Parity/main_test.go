// nolint: gocritic
package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	偶数、奇数の順にソートして返す。各要素は順不同。
	返す配列の要素数はinputと同数で、int型なので、初期化時にはその要素数がゼロ値で含まれる。
*/
func sortArrayByParity(A []int) []int {
	left, right, l := 0, len(A)-1, len(A)
	result := make([]int, l)

	for i := 0; i < l; i++ {
		if A[i]%2 == 0 {
			result[left] = A[i]
			left++
		} else {
			result[right] = A[i]
			right--
		}
	}

	return result
}

func sortArrayByParity2(A []int) []int {
	left, right := 0, len(A)-1

	for left < right {
		for left < right && A[left]%2 == 0 {
			left++
		}
		for left < right && A[right]%2 != 0 {
			right--
		}
		if left == right {
			break
		}
		A[left], A[right] = A[right], A[left]
	}

	return A
}

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		in  []int
		out [][]int
	}{
		{
			[]int{3, 1, 2, 4},
			[][]int{
				{2, 4, 3, 1},
				{2, 4, 1, 3},
				{4, 2, 1, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sortArrayByParity2(tt.in)
			isContained := false
			for _, out := range tt.out {
				isContained = reflect.DeepEqual(got, out)
			}
			if !isContained {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
