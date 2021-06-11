package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
	## summary
	swap two letters in s so the result is equal to goal. otherwise, return false.
*/
// nolint
func buddyStrings(s string, goal string) bool {
	sl, gl := len(s), len(goal)
	if sl != gl {
		return false
	}
	if strings.EqualFold(s, goal) {
		var chars [26]int
		for i := 0; i < sl; i++ {
			chars[s[i]-'a']++
		}
		for _, c := range chars {
			if c > 1 {
				return true
			}
		}
		return false
	}
	first, second := -1, -1
	for i := 0; i < sl; i++ {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}

func TestBuddyStrings(t *testing.T) {
	tests := []struct {
		inS string
		inT string
		out bool
	}{
		{
			"ab", "ba", true,
		},
		{
			"ab", "ab", false,
		},
		{
			"aa", "aa", true,
		},
		{
			"aaaaaaabc", "aaaaaaacb", true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %s", tt.inS, tt.inT), func(t *testing.T) {
			got := buddyStrings(tt.inS, tt.inT)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
