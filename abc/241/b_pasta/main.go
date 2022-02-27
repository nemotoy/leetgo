package main

import "fmt"

// i日目にはAの順番を問わず1本を選択するからmapで処理できる
func main() {
	var N, M int
	fmt.Scan(&N, &M)
	mp := map[int]int{}
	for i := 0; i < N; i++ {
		fmt.Scan(&i)
		mp[i]++
	}
	ans := "Yes"
	for i := 0; i < M; i++ {
		fmt.Scan(&i)
		if mp[i] == 0 {
			ans = "No"
			break
		}
		mp[i]--
	}
	fmt.Println(ans)
}
