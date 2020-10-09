package list

import (
	"testing"
)

/*
	## summary
*/
func isPalindrome(head *ListNode) bool {
	return false
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		in  *ListNode
		out bool
	}{
		{
			&ListNode{
				1,
				&ListNode{
					2,
					nil,
				},
			},
			false,
		},
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						2,
						&ListNode{
							1,
							nil,
						},
					},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in.Visualize(), func(t *testing.T) {
			got := isPalindrome(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
