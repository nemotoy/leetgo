package list

import (
	"reflect"
	"testing"
)

func TestListNode_Visualize(t *testing.T) {
	in := &ListNode{
		1, &ListNode{
			2, &ListNode{
				3, nil,
			},
		},
	}
	want := "1->2->3"
	got := in.Visualize()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
