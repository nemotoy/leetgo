package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	a -> 0, b -> 1
	letter value はアルファベットのindex
	numerical value は各 letter valueを連結して整数化したもの
	a~j
	ret: first + secode = targe
*/
func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return helper(firstWord)+helper(secondWord) == helper(targetWord)
}

func helper(s string) int {
	ret := 0
	for _, r := range s {
		// 桁数を上げるために*10
		// 先頭'a'は不要なので、その観点でも有用
		ret = ret*10 + int(r-'a')
	}
	return ret
}

func TestIsSumEqual(t *testing.T) {
	tests := []struct {
		in  string
		in2 string
		in3 string
		out bool
	}{
		{
			"acb", "cba", "cdb", true,
		},
		{
			"aaa", "a", "aab", false,
		},
		{
			"aaa", "a", "aaaa", true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s,%s,%s", tt.in, tt.in2, tt.in3), func(t *testing.T) {
			got := isSumEqual(tt.in, tt.in2, tt.in3)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
