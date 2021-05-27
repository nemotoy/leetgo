package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
  ## summary
  - begins with a vowel, append "ma" to the end of the word
  - begins with a not vowel, remove the first letter and append it to the end
  - Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1
*/
func toGoatLatin(sentence string) string {
	ss := strings.Split(sentence, " ")
	ret := make([]string, len(ss))
	for i, s := range ss {
		if isVowel(s[0]) {
			ret[i] = s
		} else {
			ret[i] = s[1:] + s[:1]
		}
		ret[i] += "ma" + strings.Repeat("a", i+1)
	}
	return strings.Join(ret, " ")
}

func isVowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

func TestToGoatLatin(t *testing.T) {
	tests := []struct {
		in  string
		out string
	}{
		{
			"I speak Goat Latin", "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			"The quick brown fox jumped over the lazy dog", "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := toGoatLatin(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
