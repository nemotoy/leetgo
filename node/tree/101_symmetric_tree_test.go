package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	中央値を境にシンメトリーになっているか。（Tree左辺の左辺と右辺の左辺が同一、またその逆も同一であること。）

	1. 再帰
*/
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val && isMirror(t1.Right, t2.Left) && isMirror(t1.Left, t2.Right)
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out bool
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			true,
		},
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:  2,
					Left: nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := isSymmetric(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
