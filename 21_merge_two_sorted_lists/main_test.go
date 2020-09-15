package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		in1 *ListNode
		in2 *ListNode
		out *ListNode
	}{
		{
			&ListNode{
				1, &ListNode{
					2, &ListNode{
						4, nil,
					},
				},
			},
			&ListNode{
				1, &ListNode{
					3, &ListNode{
						4, nil,
					},
				},
			},
			&ListNode{
				1, &ListNode{
					1, &ListNode{
						2, &ListNode{
							3, &ListNode{
								4, &ListNode{
									4, nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.in1, tt.in2), func(t *testing.T) {
			got := mergeTwoLists(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
