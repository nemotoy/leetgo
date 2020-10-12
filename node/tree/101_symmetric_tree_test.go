package list

import (
	"fmt"
	"reflect"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	## summary
*/
func isSymmetric(root *TreeNode) bool {
	return false
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		in  *TreeNode
		out bool
	}{}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := isSymmetric(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
