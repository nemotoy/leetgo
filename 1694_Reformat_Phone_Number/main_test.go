package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
  ## summary
*/
func reformatNumber(number string) string {
	ret := []string{}
	number = strings.ReplaceAll(strings.ReplaceAll(number, " ", ""), "-", "")
	l := len(number)
	for i := 0; i < l; {
		if l == i+4 {
			ret = append(ret, number[i:i+2])
			i += 2
		} else if i+3 <= l {
			ret = append(ret, number[i:i+3])
			i += 3
		} else {
			ret = append(ret, number[i:i+2])
			i += 2
		}
	}
	if len(ret) > 1 {
		return strings.Join(ret, "-")
	}
	return ret[0]
}

func TestReformatNumber(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"1-23-45 6", "123-456",
		},
		{
			"123 4-567", "123-45-67",
		},
		{
			"123 4-5678", "123-456-78",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := reformatNumber(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
