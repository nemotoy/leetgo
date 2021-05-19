package tree

import (
	"fmt"
	"reflect"
	"testing"
)

// 木の直径の長さ返す
// 直径は任意のノード間の長さ。ルートを通らない場合もある

// 両辺の深さの合計は解のように見えるが、ルートを通らない場合を考慮する必要があるためNG。
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(
		max(diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)),
		maxDepth(root.Left)+maxDepth(root.Right),
	)
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out int
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			3,
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
				},
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := diameterOfBinaryTree(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
