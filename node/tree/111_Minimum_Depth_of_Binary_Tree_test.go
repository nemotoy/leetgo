package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	最浅度を返す。
*/

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 検索ノードなしの場合はroot値分の1を返す。
	if root.Left == nil && root.Right == nil {
		return 1
	}
	// left/rightノードがそれぞれ存在する方を検索対象とし、再帰呼び出しする。
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}

	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}

	// 両ノードが存在する場合は、探索結果が小さい方を返す。
	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func TestMinDepth(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out int
	}{
		{
			&TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minDepth(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
