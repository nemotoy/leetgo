package tree

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

/*
	## summary
*/
func maximumUnits(boxTypes [][]int, truckSize int) int {
	// ボックスタイプをボックスあたりのユニット数で並べ替える
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	ret := 0
	for _, box := range boxTypes {
		if truckSize >= box[0] {
			ret += box[0] * box[1]
			truckSize -= box[0]
		} else {
			ret += truckSize * box[1]
			break
		}
	}
	return ret
}

func TestMaximumUnits(t *testing.T) {
	tests := []struct {
		in  [][]int
		in2 int
		out int
	}{
		{
			[][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8,
		},
		{
			[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.in), func(t *testing.T) {
			got := maximumUnits(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
