//nolint:gocritic
package main

import (
	"fmt"
	"testing"
)

/*
  ## summary
  item = [type, color, name]

  - ruleKey == "type" and ruleValue == typei.
  - ruleKey == "color" and ruleValue == colori.
  - ruleKey == "name" and ruleValue == namei.

  ruleKey is equal to either "type", "color", or "name"
*/
func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	ret := 0
	num := 0 // ruleKey is equal to "type"
	if ruleKey == "color" {
		num = 1
	} else if ruleKey == "name" {
		num = 2
	}
	for _, item := range items {
		if item[num] == ruleValue {
			ret++
		}
	}
	return ret
}

func TestCountMatches(t *testing.T) {
	tests := []struct {
		in  [][]string
		in2 string
		in3 string
		out int
	}{
		{
			[][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "lenovo"},
				{"phone", "gold", "iphone"},
			},
			"color",
			"silver",
			1,
		},
		{
			[][]string{
				{"phone", "blue", "pixel"},
				{"computer", "silver", "phone"},
				{"phone", "gold", "iphone"},
			},
			"type",
			"phone",
			2,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %s, %s", tt.in, tt.in2, tt.in3), func(t *testing.T) {
			got := countMatches(tt.in, tt.in2, tt.in3)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
