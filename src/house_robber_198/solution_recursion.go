package house_robber_198

import "fmt"

func maxRob(nums []int, cache map[int]int, idx int) int {
	if idx >= len(nums) {
		return 0
	}

	value, found := cache[idx]
	if found {
		return value
	}

	result := max(
		nums[idx]+maxRob(nums, cache, idx+2),
		maxRob(nums, cache, idx+1),
	)

	cache[idx] = result
	return result
}

func robRecursion(nums []int) int {
	cache := make(map[int]int)

	return maxRob(nums, cache, 0)
}

func TestRecursion() {
	fmt.Println(robRecursion([]int{1, 2, 3, 1}))
	fmt.Println(robRecursion([]int{2, 7, 9, 3, 1}))
}
