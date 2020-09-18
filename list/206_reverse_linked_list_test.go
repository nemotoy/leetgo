package list

import (
	"reflect"
	"testing"
)

/*
	## summary
*/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
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
				if got != nil && tt.out != nil {
					t.Errorf("details; got: %s, want: %s", got.Visualize(), tt.out.Visualize())
				} else {
					t.Errorf("got: %v, want: %v", got, tt.out)
				}
			}
		})
	}
}
