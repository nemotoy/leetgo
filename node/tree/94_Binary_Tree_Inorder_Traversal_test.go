package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  間順走査し、値を順番通り返す。
  間順はWikipediaより引用。
  1. もしあれば、左の部分木を間順走査する。
  2. 根ノードを調査する。
  3. もしあれば、右の部分木を間順走査する。
*/
func inorderTraversal(root *TreeNode) []int {
	ret := []int{}
	helperInorderTraversal(root, &ret)
	return ret
}

func helperInorderTraversal(root *TreeNode, ret *[]int) {
	if root != nil {
		if root.Left != nil {
			helperInorderTraversal(root.Left, ret)
		}
		*ret = append(*ret, root.Val)
		if root.Right != nil {
			helperInorderTraversal(root.Right, ret)
		}
	}
}

func TestInorderTraversal(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out []int
	}{
		{
			&TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			[]int{1, 3, 2},
		},
		{
			nil,
			[]int{},
		},
		{
			&TreeNode{
				Val: 1,
			},
			[]int{1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := inorderTraversal(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
