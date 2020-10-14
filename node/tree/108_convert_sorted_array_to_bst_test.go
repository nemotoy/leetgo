package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	BST(Binary Search Tree)を返す。

	1. 配列の中央をdepth:0の値とし、その左を左辺、右を右辺としてTreeNodeを生成する。
*/
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	if l == 0 {
		return nil
	}
	mid := l / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		in  []int
		out *TreeNode
	}{
		{
			[]int{-10, -3, 0, 5, 9},
			&TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: -3,
					Left: &TreeNode{
						Val: -10,
					},
				},
				Right: &TreeNode{
					Val: 9,
					Left: &TreeNode{
						Val: 5,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := sortedArrayToBST(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
