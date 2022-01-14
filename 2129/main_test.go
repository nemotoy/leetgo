package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
	- 2以下の文字数はlowercase
	- それ以外は1文字目はuppercase,その他はlowercase
*/
func capitalizeTitle(title string) string {
	ss := strings.Split(title, " ")
	ret := ""
	for _, s := range ss {
		s = strings.ToLower(s)
		if len(s) > 2 {
			s = strings.ToUpper(string(s[:1])) + string(s[1:])
		}
		ret += s + " "
	}
	return ret[:len(ret)-1]
}

func TestCapitalizeTitle(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"capiTalIze tHe titLe", "Capitalize The Title",
		},
		{
			"First leTTeR of EACH Word", "First Letter of Each Word",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := capitalizeTitle(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
