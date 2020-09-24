package main

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	## summary
	素数評価 正の約数が1と自身
*/
func countPrimes(n int) int {
	r := 0
	for i := 1; i < n; i++ {
		if isPrime(i) {
			r++
		}
	}
	// 自身は評価しない。基底値-1 から2以上の間で割れるか。

	return r
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			10, 4,
		},
		{
			2, 0,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.in), func(t *testing.T) {
			got := countPrimes(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
