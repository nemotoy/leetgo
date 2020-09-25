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
func countPrimes_old(n int) int {
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
	1. 2からxまでの整数の昇順リストを作成する（以下、探索リスト）
	2. 探索リストの先頭の要素を素数リスト（最終的に返すのはこのリストの要素数）に移動する。探索リストからその要素の倍数を削除する
	3. 2の操作を、探索リストの先頭値がxの平方根に達するまで行う
	4. 探索リストに残った数を素数リストに移動して処理終了
*/
func countPrimesSoE(n int) int {

	if n <= 2 {
		return 0
	}

	searchList := make([]int, n-2)
	for i := 0; i+2 < n; i++ {
		searchList[i] = i + 2
	}
	sqrt := math.Sqrt(float64(n))
	primeList := []int{}
	for len(searchList) > 0 {
		fn := searchList[0]
		primeList = append(primeList, fn)
		if float64(fn) > sqrt {
			primeList = append(primeList, searchList[1:]...)
			break
		}
		searchList = remove(searchList, fn)
	}
	return len(primeList)
}

func remove(nums []int, n int) []int {
	r := []int{}
	for _, v := range nums {
		if v%n != 0 {
			r = append(r, v)
		}
	}
	return r
}

func countPrimes(n int) int {
	primeList := make([]bool, n)
	for i := 2; i < n; i++ {
		primeList[i] = true
	}

	// iで割り切れれば、i*iもiで割り切れる。
	for i := 2; i*i < n; i++ {
		if !primeList[i] {
			continue
		}
		for j := i * i; j < n; j += i {
			primeList[j] = false
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if primeList[i] {
			count++
		}
	}
	return count
}

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		in  int
		out int
	}{
		{
			10, 4, // 3,5,7,9 = 4
		},
		{
			2, 0,
		},
		{
			4, 2,
		},
		{
			3, 1,
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
