package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	与えられたTreeNodeがBSTかどうかを評価する。

*/
func isValidBST(root *TreeNode) bool {
	return false
}

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out bool
	}{
		{
			&TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			true,
		},
		{
			&TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 6,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := isValidBST(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
