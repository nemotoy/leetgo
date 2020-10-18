package main

import (
	"fmt"
	"testing"
	"unicode"
)

/*
	## summary
	lower letterとdigitsを隣接させる。
*/
func reformat(s string) string {
	var let, dig string
	for _, r := range s {
		if unicode.IsNumber(r) {
			dig += string(r)
		} else {
			let += string(r)
		}
	}
	if let == "" || dig == "" {
		return ""
	}
	fmt.Println(let, dig)
	res := ""
	if len(let) > len(dig) {
		let, dig = dig, let
	}
	for i := 0; i < len(s)/2; i++ {
		res += string(dig[i])
		res += string(let[i])
		fmt.Println(res)
	}
	return res
}

func TestReformat(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"a0b1c2", "0a1b2c",
		},
		{
			"1229857369", "",
		},
		{
			"covid2019", "c2o0v1i9d",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := reformat(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
