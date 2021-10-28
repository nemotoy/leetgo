package main

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

/*
 - ++X and X++ increments the value of the variable X by 1.
 - --X and X-- decrements the value of the variable X by 1.

 return the final value of X after performing all the operations.
*/
func finalValueAfterOperations(operations []string) int {
	x := 0
	for _, op := range operations {
		if op == "X++" || op == "++X" {
			x++
		} else {
			x--
		}
	}
	return x
}

func finalValueAfterOperations2(operations []string) int {
	x := 0
	for _, op := range operations {
		if strings.Contains(op, "++") {
			x++
		} else {
			x--
		}
	}
	return x
}

func TestFinalValueAfterOperations(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{[]string{"--X", "X++", "X++"}, 1},
		{[]string{"++X", "++X", "X++"}, 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := finalValueAfterOperations2(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
