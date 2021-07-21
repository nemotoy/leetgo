package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
	Given an array of strings words, return the words that can be typed using letters of the alphabet on only one row of American keyboard like the image below.
*/
func findWords(words []string) []string {
	first, second, third := "qwertyuiop", "asdfghjkl", "zxcvbnm"
	keyboards := []string{first, second, third}
	keyMaps := make(map[rune]int, len(first+second+third))
	for row, chars := range keyboards {
		for _, char := range chars {
			keyMaps[char] = row
		}
	}

	ret := make([]string, 0, len(words))
	for _, word := range words {
		lower := strings.ToLower(word)
		row := keyMaps[rune(lower[0])]
		onlyRow := true
		for _, char := range lower[1:] {
			if keyMaps[char] != row {
				onlyRow = false
				break
			}
		}
		if onlyRow {
			ret = append(ret, word)
		}
	}
	return ret
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		in  []string
		out []string
	}{
		{
			[]string{"Hello", "Alaska", "Dad", "Peace"},
			[]string{"Alaska", "Dad"},
		},
		{
			[]string{"omk"},
			[]string{},
		},
		{
			[]string{"adsdf", "sfd"},
			[]string{"adsdf", "sfd"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findWords(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
