package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
*/
func lemonadeChange(bills []int) bool {
	five, ten := 0, 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}
			five--
			ten++
		case 20:
			switch {
			case five > 0 && ten > 0:
				five--
				ten--
			case five >= 3:
				five -= 3
			default:
				return false
			}
		}
	}
	return true
}

func TestLemonadeChange(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{[]int{5, 5, 5, 10, 20}, true},
		{[]int{5, 5, 10, 10, 20}, false},
		{[]int{5, 5, 10}, true},
		{[]int{10, 10}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := lemonadeChange(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
