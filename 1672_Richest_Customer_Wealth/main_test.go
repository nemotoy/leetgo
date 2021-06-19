package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	各要素の最大値
*/
func maximumWealth(accounts [][]int) int {
	ret := 0
	for _, as := range accounts {
		tmp := 0
		for _, a := range as {
			tmp += a
		}
		if ret < tmp {
			ret = tmp
		}
	}
	return ret
}

func TestMaximumWealth(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{1, 2, 3}, {3, 2, 1}}, 6,
		},
		{
			[][]int{{1, 5}, {7, 3}, {3, 5}}, 10,
		},
		{
			[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}, 17,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maximumWealth(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
