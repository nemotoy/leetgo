// nolint:gocritic
package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	kにいるpersonがticketを購入し終わるまでの時間
*/
func timeRequiredToBuy(tickets []int, k int) int {
	time := 0
	for len(tickets) > 0 {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] == 0 {
				continue
			}
			tickets[i]--
			time++
			if tickets[k] == 0 {
				return time
			}
		}
	}
	return time
}

// 列のチケット数より少ない場合は、indexのチケット数購入する
// 上記以外で、列より前の場合は、列のチケット数分購入。列より後ろの場合は、列のチケット数より1枚少なく購入
func timeRequiredToBuy2(tickets []int, k int) int {
	time := 0
	ticket := tickets[k]
	for i := 0; i < len(tickets); i++ {
		if tickets[i] < ticket {
			time += tickets[i]
		} else if i <= k {
			time += ticket
		} else if i > k {
			time += ticket - 1
		}
	}
	return time
}

func TestTimeRequiredToBuy(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out int
	}{
		{[]int{2, 3, 2}, 2, 6},
		{[]int{5, 1, 1, 1}, 0, 8},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := timeRequiredToBuy(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
