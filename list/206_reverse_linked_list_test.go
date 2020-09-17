package list

import (
	"reflect"
	"testing"
)

/*
	## summary
*/
func reverseList(head *ListNode) *ListNode {
	r := &ListNode{}
	for head.Next != nil {
		r.Next = head
		head = head.Next
	}
	return r.Next
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		in  *ListNode
		out *ListNode
	}{
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
			&ListNode{
				5,
				&ListNode{
					4,
					&ListNode{
						3,
						&ListNode{
							2,
							&ListNode{
								1,
								nil,
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.in.Visualize(), func(t *testing.T) {
			got := reverseList(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
				if got != nil && tt.out != nil {
					t.Logf("details; got: %v, want: %v", got.Visualize(), tt.out.Visualize())
				}
			}
		})
	}
}
