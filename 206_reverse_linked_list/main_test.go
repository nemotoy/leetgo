package main

import (
	"reflect"
	"testing"

	l "nemotoy/leetgo/list"
)

/*
	## summary
*/
func reverseList(head *l.ListNode) *l.ListNode {
	r := &l.ListNode{}
	for head.Next != nil {
		r.Next = head
		head = head.Next
	}
	return r.Next
}

func TestReverseList(t *testing.T) {
	tests := []struct {
		in  *l.ListNode
		out *l.ListNode
	}{
		{
			&l.ListNode{
				Val: 1,
				Next: &l.ListNode{
					Val: 2,
					Next: &l.ListNode{
						Val: 3,
						Next: &l.ListNode{
							Val: 4,
							Next: &l.ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			&l.ListNode{
				Val: 5,
				Next: &l.ListNode{
					Val: 4,
					Next: &l.ListNode{
						Val: 3,
						Next: &l.ListNode{
							Val: 2,
							Next: &l.ListNode{
								Val:  1,
								Next: nil,
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
