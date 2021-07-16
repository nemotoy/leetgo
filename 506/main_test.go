package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
	## summary
	score[i] = rank
*/
func findRelativeRanks(score []int) []string {
	l := len(score)
	ret := make([]string, 0, l)
	for i := 0; i < l; i++ {
		n := score[i]
		place := 0
		for j := 0; j < l; j++ {
			if n > score[j] {
				place++
			}
		}
		rank := ""
		switch place {
		case l - 1:
			rank = "Gold Medal"
		case l - 2:
			rank = "Silver Medal"
		case l - 3:
			rank = "Bronze Medal"
		default:
			rank = strconv.Itoa(l - place)
		}
		ret = append(ret, rank)
	}
	return ret
}

func TestFindRelativeRanks(t *testing.T) {
	tests := []struct {
		in  []int
		out []string
	}{
		{
			[]int{5, 4, 3, 2, 1},
			[]string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"},
		},
		{
			[]int{10, 3, 8, 9, 4},
			[]string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := findRelativeRanks(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
