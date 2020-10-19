package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	@前:
	- "alice.z@leetcode.com" equals to  "alicez@leetcode.com"
	- "m.y+name@email.com" equals to "my@email.com"
	@後: 同値。ルール適用外。

	"."は無視される。
	"+"はそれ以降が無視される。

	".", "+"を整理する。
*/
func numUniqueEmails(emails []string) int {
	hashTable := make(map[string]struct{}, len(emails))
	for _, email := range emails {
		for i := 0; i < len(email); i++ {
			e := string(email[i])
			if e == "@" {
				// add email to hashTable
				fmt.Println(email)
				hashTable[email] = struct{}{}
				break
			}
			if e == "." {
				// remove "." element
				email = email[:i] + email[i+1:]
			}
			if e == "+" {
				// remove elements from "+" to "@"
				for string(email[i]) != "@" {
					email = email[:i] + email[i+1:]
				}
				hashTable[email] = struct{}{}
				break
			}
		}
	}
	return len(hashTable)
}

func TestNumUniqueEmails(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{
			[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := numUniqueEmails(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
