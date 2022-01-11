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
