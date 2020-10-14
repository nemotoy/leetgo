package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary


*/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func TestMaxDepth(t *testing.T) {
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
			3,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maxDepth(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
