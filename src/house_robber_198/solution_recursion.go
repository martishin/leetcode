package house_robber_198

import "fmt"

func maxRob(nums []int, idx int, cache []int) int {
	if idx >= len(nums) {
		return 0
	}

	found := cache[idx]
	if found != -1 {
		return found
	}

	result := max(
		nums[idx]+maxRob(nums, idx+2, cache),
		maxRob(nums, idx+1, cache),
	)

	cache[idx] = result
	return result
}

func robRecursion(nums []int) int {
	cache := make([]int, len(nums))
	for i := range len(cache) {
		cache[i] = -1
	}

	return maxRob(nums, 0, cache)
}

func TestRecursion() {
	fmt.Println(robRecursion([]int{1, 2, 3, 1}))
	fmt.Println(robRecursion([]int{2, 7, 9, 3, 1}))
}
