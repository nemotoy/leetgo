package tree

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	piecesからarrを生成可能か。ただしpieces[i]は並べ替え不可。
*/
func canFormArray(arr []int, pieces [][]int) bool {
	firstEles := make(map[int][]int, len(pieces))
	for i := 0; i < len(pieces); i++ {
		firstEles[pieces[i][0]] = pieces[i]
	}

	i := 0
	for i < len(arr) {
		list, ok := firstEles[arr[i]]
		if !ok {
			return false
		}
		for j := 0; j < len(list); j++ {
			if arr[i] != list[j] {
				return false
			}
			i++
		}
	}

	return true
}

func TestCanFormArray(t *testing.T) {
	tests := []struct {
		in  []int
		in2 [][]int
		out bool
	}{
		{
			[]int{85}, [][]int{{85}}, true,
		},
		{
			[]int{15, 88}, [][]int{{88}, {15}}, true,
		},
		{
			[]int{49, 18, 16}, [][]int{{16, 18, 49}}, false,
		},
		{
			[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}, true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.in, tt.in2), func(t *testing.T) {
			got := canFormArray(tt.in, tt.in2)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
