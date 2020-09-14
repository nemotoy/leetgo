package main

import (
	"reflect"
	"testing"
)

/*
	## summary
	エクセルのような列番号（項目）を取得する。
	1~26: A ~ Z
	27~52: AA ~ AZ
	53~  : BA
*/
func titleToNumber(s string) int {
	return 0
}

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		in  string
		out int
	}{
		{
			"A", 1,
		},
		{
			"AA", 27,
		},
		{
			"AB", 28,
		},
		{
			"ZY", 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := titleToNumber(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
