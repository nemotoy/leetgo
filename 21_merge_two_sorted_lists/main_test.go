package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
	## summary
	2つの昇順ListNodeをマージして1つの昇順リストにする。要素の重複あり。

	- どのように探索するか
		- 数字の抽出。マージ。昇順ソート。ListNode生成。
		- 両listの1つ目を比較して、小さい方を付ける。以降繰り返す。
	- Goでどう書くか
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	result := &ListNode{}
	tmp := result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			tmp.Next = l2
			l2 = nil
			continue
		}
		if l2 == nil {
			tmp.Next = l1
			l1 = nil
			continue
		}
		if l1.Val < l2.Val {
			tmp.Next = l1
			tmp = tmp.Next
			l1 = l1.Next
		} else {
			tmp.Next = l2
			tmp = tmp.Next
			l2 = l2.Next
		}
	}
	return result.Next
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
		t.Run(fmt.Sprintf("%s, %s", chainedLists(tt.in1, ""), chainedLists(tt.in2, "")), func(t *testing.T) {
			got := mergeTwoLists(tt.in1, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
				if got != nil && tt.out != nil {
					t.Logf("details; got: %v, want: %v", chainedLists(got, ""), chainedLists(tt.out, ""))
				}
			}
		})
	}
}

const listSurfix = "->"

// 1->2->3 のような文字列を返す
func chainedLists(l *ListNode, s string) string {
	v := strconv.Itoa(l.Val)
	if l.Next == nil {
		return s + v
	}
	return chainedLists(l.Next, s+v+listSurfix)
}

func TestChainedLists(t *testing.T) {
	in := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, nil,
			},
		},
	}
	want := "1->2->3"
	got := chainedLists(in, "")
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
