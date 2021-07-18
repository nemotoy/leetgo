package main

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

/*
	## summary
	+ -> sum of the previous two scores
	D -> double the previous scores
	C -> remove previous score from the record
*/
func calPoints(ops []string) int {
	records := make([]int, 0, len(ops))
	for _, op := range ops {
		switch op {
		case "+":
			l := len(records)
			records = append(records, records[l-1]+records[l-2])
		case "D":
			records = append(records, records[len(records)-1]*2)
		case "C":
			records = records[:len(records)-1]
		default:
			v, _ := strconv.Atoi(op)
			records = append(records, v)
		}
	}
	ret := 0
	for _, r := range records {
		ret += r
	}
	return ret
}

func TestCalPoints(t *testing.T) {
	tests := []struct {
		in  []string
		out int
	}{
		{
			[]string{"5", "2", "C", "D", "+"}, 30,
		},
		{
			[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}, 27,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := calPoints(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
