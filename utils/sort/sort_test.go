package sort

import (
	"fmt"
	"reflect"
	"testing"
)

// 挿入ソート
func InsertionSort(a []int) {
	n := len(a)
	// 1要素目はそのまま
	for i := 1; i < n; i++ {
		// 挿入する値
		v := a[i]
		j := i
		for ; j > 0; j-- {
			// vよりも大きい要素は1つ後ろに移す
			if a[j-1] > v {
				a[j] = a[j-1]
			} else {
				break
			}
		}
		// j番目にvを挿入
		a[j] = v
	}
}

// スライスaの区間left,rightをソートする
func MergeSort(a []int, left, right int) {

}

func QuickSort(a []int, left, right int) {
	if right-left <= 1 {
		return
	}
	// pivotを適当に中点とする
	pivotIndex := (left + right) / 2
	pivot := a[pivotIndex]
	// pivotと右端をswap
	swap(a, pivotIndex, right-1)

	i := left
	for j := left; j < right-1; j++ {
		// pivot未満のものは左に詰める
		if a[j] < pivot {
			swap(a, i, j)
			i++
		}
	}
	// pivotを適切な場所に挿入
	swap(a, i, right-1)

	QuickSort(a, left, i)
	QuickSort(a, i+1, right)
}

func BucketSort(a []int) {

}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{4, 1, 3, 5, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 1, 3, 5, 0}, []int{0, 0, 1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			InsertionSort(tt.in)
			if !reflect.DeepEqual(tt.in, tt.out) {
				t.Errorf("got: %v, want: %v", tt.in, tt.out)
			}
		})
	}
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{4, 1, 3, 5, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 1, 3, 5, 0}, []int{0, 0, 1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			QuickSort(tt.in, 0, len(tt.in))
			if !reflect.DeepEqual(tt.in, tt.out) {
				t.Errorf("got: %v, want: %v", tt.in, tt.out)
			}
		})
	}
}

func TestBucketSort(t *testing.T) {
	tests := []struct {
		in  []int
		out []int
	}{
		{[]int{4, 1, 3, 5, 2}, []int{1, 2, 3, 4, 5}},
		{[]int{0, 1, 3, 5, 0}, []int{0, 0, 1, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			BucketSort(tt.in)
			if !reflect.DeepEqual(tt.in, tt.out) {
				t.Errorf("got: %v, want: %v", tt.in, tt.out)
			}
		})
	}
}
