package main

import (
	l "nemotoy/leetgo/list"
	"reflect"
	"testing"
)

/*
	## summary
	2つの片方向リストをマージして1つの片方向リストにする。各ノードの要素はint型で、それは昇順になる。また、重複は許容する。
*/
func mergeTwoLists(l1 *l.ListNode, l2 *l.ListNode) *l.ListNode {

	result := &l.ListNode{}
	// tailの値はListNodeオブジェクトのポインタを持ったresultで、ある時点以降のListNodeとして扱う。
	for tail := result; l1 != nil || l2 != nil; tail = tail.Next {
		// 片方がnilになれば、一方のlistNodeを末尾に追加するだけで良いためイテレーションから抜ける。
		if l1 == nil {
			tail.Next = l2
			break
		}
		if l2 == nil {
			tail.Next = l1
			break
		}
		if l1.Val < l2.Val {
			// 末尾に小さい数を持つListNodeを追加する。
			tail.Next = l1
			// 追加されたListNodeは、次点に進める。
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
	}
	return result.Next
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		in1 *l.ListNode
		in2 *l.ListNode
		out *l.ListNode
	}{
		{
			&l.ListNode{
				Val: 1,
				Next: &l.ListNode{
					Val: 2,
					Next: &l.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&l.ListNode{
				Val: 1,
				Next: &l.ListNode{
					Val: 3,
					Next: &l.ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&l.ListNode{
				Val: 1,
				Next: &l.ListNode{
					Val: 1,
					Next: &l.ListNode{
						Val: 2,
						Next: &l.ListNode{
							Val: 3,
							Next: &l.ListNode{
								Val: 4,
								Next: &l.ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.in1.Visualize()+tt.in2.Visualize(), func(t *testing.T) {
			got := mergeTwoLists(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
				if got != nil && tt.out != nil {
					t.Logf("details; got: %s, want: %s", got.Visualize(), tt.out.Visualize())
				}
			}
		})
	}
}
