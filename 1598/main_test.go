package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
  フォルダの変更操作後にメインフォルダに戻るために必要な最小操作数を返す。
*/
func minOperations(logs []string) int {
	dir := 0
	for i := 0; i < len(logs); i++ {
		switch logs[i] {
		case "../":
			if dir > 0 {
				dir--
			}
		case "./":
		default:
			dir++
		}
	}
	return dir
}

func TestMinOperations(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{[]string{"d1/", "d2/", "../", "d21/", "./"}, 2},
		{[]string{"d1/", "d2/", "./", "d3/", "../", "d31/"}, 3},
		{[]string{"d1/", "../", "../", "../"}, 0},
		{[]string{"d1/", "d2/", "../", "d21/", "./"}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := minOperations(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
