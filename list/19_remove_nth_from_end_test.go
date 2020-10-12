package list

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	連結リストの先頭からn要素目を削除する。
	連結リストの要素数をLとすると、削除する要素は L - n+1 となる。ここでのLはリストの次要素がnilである場合を終端とするまでの長さ。
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	l := 0
	first := head

	for first != nil {
		l++
		first = first.Next
	}
	l -= n
	first = dummy
	for l > 0 {
		l--
		first = first.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		in1 *ListNode
		in2 int
		out *ListNode
	}{
		{
			&ListNode{
				1, &ListNode{
					2, &ListNode{
						3, &ListNode{
							4, &ListNode{
								5, nil,
							},
						},
					},
				},
			},
			2,
			&ListNode{
				1, &ListNode{
					2, &ListNode{
						3, &ListNode{
							5, nil,
						},
					},
				},
			},
		},
		{
			&ListNode{
				1, nil,
			},
			1,
			nil,
		},
		{
			&ListNode{
				1, &ListNode{
					2, nil,
				},
			},
			1,
			&ListNode{
				1, nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("rm(%d)%s", tt.in2, tt.in1.Visualize()), func(t *testing.T) {
			got := removeNthFromEnd(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
				if got != nil && tt.out != nil {
					t.Logf("details; got: %s, want: %s", got.Visualize(), tt.out.Visualize())
				}
			}
		})
	}
}
