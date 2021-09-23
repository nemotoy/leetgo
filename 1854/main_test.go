package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
 */
func maximumPopulation(logs [][]int) int {
	pop, res := [2051]int{}, 0
	for _, log := range logs {
		pop[log[0]]++
		pop[log[1]]--
	}
	for i := 1950; i < 2050; i++ {
		pop[i] += pop[i-1]
		if pop[i] > pop[res] {
			res = i
		}
	}
	return res
}

func TestMaximumPopulation(t *testing.T) {
	tests := []struct {
		in  [][]int
		out int
	}{
		{
			[][]int{{1993, 1999}, {2000, 2010}}, 1993,
		},
		{
			[][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}, 1960,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maximumPopulation(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
