package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary（two-pointer approach）
	- バイト配列の両端を始点としたint型を定義する（l,r）
	- ポインターが指す値を入れ替える。先頭が末尾を超過するまでループする。
*/
func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l, r = l+1, r-1
	}
}

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		in  []byte
		out []byte
	}{
		{
			[]byte("hello"),
			[]byte("olleh"),
		},
		{
			[]byte("Hannah"),
			[]byte("hannaH"),
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s", tt.in), func(t *testing.T) {
			reverseString(tt.in)
			got := tt.in
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
