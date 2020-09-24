package main

import (
	"fmt"
	"math"
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

/*
	エラトステネスの篩（Sieve of Eratosthenes）：指定された整数以下の全ての素数を発見するための単純なアルゴリズム
*/
func countPrimesWithSoE(n int) int {

	searchList := []int{}
	for i := 2; i <= n; i++ {
		if i%2 != 0 {
			searchList = append(searchList, i)
		}
	}
	fmt.Println(searchList)
	if len(searchList) == 0 {
		return 0
	}

	primeList := []int{searchList[0]}

	fmt.Println(primeList)

	sqrt := math.Sqrt(float64(n))
	fmt.Println(sqrt)
	for i, n1 := range searchList {
		if float64(n1) > sqrt {
			fmt.Println(primeList, searchList[i:])
			primeList = append(primeList, searchList[i:]...)
			break
		}
		primeList = append(primeList, n)
	}

	return len(searchList)
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
			got := countPrimesWithSoE(tt.in)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got: %v, want: %v", got, tt.out)
			}
		})
	}
}
