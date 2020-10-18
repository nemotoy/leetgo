package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	配列の最後尾が接頭に該当しなかったら、それを返す。
	- destがsourceにない、destを返す。
*/
func destCity(paths [][]string) string {
	p2 := make(map[string]string, len(paths))
	for _, p := range paths {
		p2[p[0]] = p[1]
	}
	for _, d := range p2 {
		if _, ok := p2[d]; !ok {
			return d
		}
	}
	return ""
}

func TestDestCity(t *testing.T) {
	tests := []struct {
		in  [][]string
		out string
	}{
		{
			[][]string{
				{"London", "New York"},
				{"New York", "Lima"},
				{"Lima", "Sao Paulo"},
			},
			"Sao Paulo",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := destCity(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
