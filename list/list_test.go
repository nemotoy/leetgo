package list

import (
	"reflect"
	"testing"
)

func TestListNode_Visualize(t *testing.T) {
	tests := []struct {
		in  *ListNode
		out string
	}{
		{
			&ListNode{
				1, &ListNode{
					2, &ListNode{
						3, nil,
					},
				},
			},
			"1->2->3",
		},
		{
			nil,
			"",
		},
	}
	for _, tt := range tests {
		got := tt.in.Visualize()
		if !reflect.DeepEqual(got, tt.out) {
			t.Errorf("got: %s, want: %s", got, tt.out)
		}
	}
}
