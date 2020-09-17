package list

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

const nodeLink = "->"

// ListNodeオブジェクトから文字列（1->2->3）を返す
func (l *ListNode) Visualize() string {
	s := ""
	for l != nil {
		v := strconv.Itoa(l.Val)
		if l.Next == nil {
			s += v
			break
		}
		s += v + nodeLink
		l = l.Next
	}
	return s
}
