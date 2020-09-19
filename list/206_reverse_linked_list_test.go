package list

import (
	"reflect"
	"testing"
)

/*
	## summary

	- ログ
	head: 1->2->3->4->5, next: 2->3->4->5
	head: 2->3->4->5, next: 3->4->5
	head: 3->4->5, next: 4->5
	head: 4->5, next: 5
	head: 5
	head: 4->5, p: 5
	head: 3->4, p: 5->4
	head: 2->3, p: 5->4->3
	head: 1->2, p: 5->4->3->2
*/
func reverseList(head *ListNode) *ListNode {
	// 末尾要素まで探索
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next) // headを返すまではここで再帰する。
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
