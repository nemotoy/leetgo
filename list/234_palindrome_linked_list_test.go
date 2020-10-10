package list

import (
	"testing"
)

/*
	## summary
	struct型の場合、string型と違い長さがわからない（接尾から探索できない。）

	- Nextがnilになるまで探索して、要素をメモリに保存する（int型のスライス）
	- スライスを操作する
*/
func isPalindrome(head *ListNode) bool {
	nums := []int{}

	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
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
