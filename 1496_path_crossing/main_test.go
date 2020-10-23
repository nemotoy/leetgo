package main

import (
	"testing"
)

/*
	## summary
	NESWは方位を指す。以前通過したポイントを止まる場合は真を返す。
	1. 各座標を持つデータ構造を実装し、毎回参照して、同じものが含まれていれば真を返すようにする。
	- データ構造
	  - (1,0), (1,1)が取り得る
*/
func isPathCrossing(path string) bool {
	type coord [2]int
	paths := map[coord]bool{
		{0, 0}: true,
	}
	var ns, ew int // (x, y)
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'N':
			ns++
		case 'E':
			ew++
		case 'S':
			ns--
		case 'W':
			ew--
		}
		if _, ok := paths[coord{ns, ew}]; ok {
			return true
		}
		paths[coord{ns, ew}] = true
	}
	return false
}

func TestIsPathCrossing(t *testing.T) {
	tests := []struct {
		in  string
		out bool
	}{
		{
			"NES", false,
		},
		{
			"NESWW", true,
		},
		{
			"NNSWWEWSSESSWENNW", true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got := isPathCrossing(tt.in)
			if got != tt.out {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
