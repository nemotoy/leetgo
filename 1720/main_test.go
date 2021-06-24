package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func decode(encoded []int, first int) []int {
	l := len(encoded)
	ret := make([]int, l+1)
	ret[0] = first
	for i := 0; i < l; i++ {
		ret[i+1] = ret[i] ^ encoded[i]
	}
	return ret
}

func TestDecode(t *testing.T) {
	tests := []struct {
		in  []int
		in2 int
		out []int
	}{
		{
			[]int{1, 2, 3},
			1,
			[]int{1, 0, 2, 1},
		},
		{
			[]int{6, 2, 7, 3},
			4,
			[]int{4, 2, 0, 7, 4},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.in, tt.in2), func(t *testing.T) {
			got := decode(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
