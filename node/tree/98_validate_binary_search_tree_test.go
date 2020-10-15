package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	与えられたTreeNodeがBSTかどうかを評価する。等号は含まれない。
*/
func isValidBST(root *TreeNode) bool {
	return helper(root, nil, nil)
}

// 型シグネチャのlower/upperには 0 も含まれるので、ポインタ型で定義する。
// 再帰的に左辺・右辺のサブツリーを探索する。nodeがnilになる場合に真を返す。
func helper(node *TreeNode, lower, upper *int) bool {
	if node == nil {
		return true
	}
	val := node.Val
	if lower != nil && val <= *lower {
		return false
	}
	if upper != nil && val >= *upper {
		return false
	}
	return helper(node.Left, lower, &val) && helper(node.Right, &val, upper)
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
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 1,
				},
			},
			false,
		},
		{
			&TreeNode{
				Val: 10,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 15,
					Left: &TreeNode{
						Val: 6,
					},
					Right: &TreeNode{
						Val: 20,
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
