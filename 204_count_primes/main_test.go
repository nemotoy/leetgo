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
	1. 2からxまでの整数の昇順リストを作成する（以下、探索リスト）
	2. 探索リストの先頭の要素を素数リスト（最終的に返すのはこのリストの要素数）に移動する。探索リストからその要素の倍数を削除する
	3. 2の操作を、探索リストの先頭値がxの平方根に達するまで行う
	4. 探索リストに残った数を素数リストに移動して処理終了
*/
func countPrimesWithSoE(n int) int {

	searchList := make([]int, n-1)
	for i := 0; i+2 <= n; i++ {
		searchList[i] = i + 2
	}
	sqrt := math.Sqrt(float64(n))

	fmt.Println(searchList)
	primeList := []int{}
	for len(searchList) > 0 {
		fn := searchList[0]
		if float64(fn) > sqrt {
			fmt.Println("to sqrt")
			break
		}
		primeList = append(primeList, fn) //todo
		for i, s := range searchList {
			if s%fn == 0 {
				fmt.Println(s, fn)
				searchList = append(searchList[:i], searchList[i+1:]...)
				fmt.Println(searchList)
			}
		}
	}
	return len(primeList)
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
		{
			4, 2,
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
