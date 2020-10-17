package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func freqAlphabets(s string) string {
	r := ""
	l := len(s)
	for i := 0; i < l; {
		if i+2 < l && string(s[i+2]) == "#" {
			ss := (s[i]-'0')*10 + (s[i+1] - '0')
			r += string(ss + 'a' - 1)
			i += 3
		} else {
			r += string(s[i] - '0' - 1 + 'a')
			i++
		}
	}
	return r
}

func TestFreqAlphabets(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"10#11#12", "jkab",
		},
		{
			"1326#", "acz",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := freqAlphabets(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
