package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &nums[i])
	}
	sort.Ints(nums)
	fmt.Println(nums[0], nums[1], nums[2])
}
