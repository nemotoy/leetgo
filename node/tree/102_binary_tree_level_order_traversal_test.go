package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	同じ深度の要素を配列に見立てる。
*/
func levelOrder(root *TreeNode) [][]int {
	var r [][]int
	levelOrderHelper(root, &r, 0)
	return r
}

func levelOrderHelper(node *TreeNode, r *[][]int, depth int) {
	if node == nil {
		return
	}
	if depth >= len(*r) {
		*r = append(*r, []int{})
	}
	(*r)[depth] = append((*r)[depth], node.Val)
	levelOrderHelper(node.Left, r, depth+1)
	levelOrderHelper(node.Right, r, depth+1)
}

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out [][]int
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
			[][]int{
				{3},
				{9, 20},
				{15, 7},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := levelOrder(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
