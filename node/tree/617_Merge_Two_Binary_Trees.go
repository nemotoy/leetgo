package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 木1を元に考える
	// rootを加算
	root1.Val += root2.Val
	// 両木が同一rootを加算
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}

func mergeTrees2(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  mergeTrees(root1.Left, root2.Left),
		Right: mergeTrees(root1.Right, root2.Right),
	}
}

func TestMergeTrees(t *testing.T) {
	tests := []struct {
		in1 *TreeNode
		in2 *TreeNode
		out *TreeNode
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
			&TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
					Left: &TreeNode{
						Val: 5,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 5,
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in1), func(t *testing.T) {
			got := mergeTrees(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
