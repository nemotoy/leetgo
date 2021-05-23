package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	以下の規則に従ってパースする。
	G -> G
	() -> o
	(al) -> al
*/
func interpret(command string) string {
	ret, i, l := "", 0, len(command)-1
	for i <= l {
		switch string(command[i]) {
		case "G":
			ret += "G"
			i++
		case "(":
			ss := string(command[i+1])
			if ss == ")" {
				ret += "o"
				i += 2
			} else {
				ret += "al"
				i += 4
			}
		}
	}
	return ret
}

func TestInterpret(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"G()(al)", "Goal",
		},
		{
			"G()()()()(al)", "Gooooal",
		},
		{
			"(al)G(al)()()G", "alGalooG",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := interpret(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
