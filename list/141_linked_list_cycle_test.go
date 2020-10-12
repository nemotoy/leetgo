package list

import (
	"testing"
)

/*
	## summary
*/
func hasCycle(head *ListNode) bool {
	return false
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		in  *ListNode
		out bool
	}{
		{
			&ListNode{
				3,
				&ListNode{
					2,
					&ListNode{
						0,
						&ListNode{
							-4, nil,
						},
					},
				},
			},
			true,
		},
		{
			&ListNode{
				1,
				&ListNode{
					2, nil,
				},
			},
			true,
		},
		{
			&ListNode{
				1, nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in.Visualize(), func(t *testing.T) {
			got := hasCycle(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
