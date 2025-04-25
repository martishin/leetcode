package subsets_78

import "fmt"

func backtrack(nums []int, idx int, currentNums []int, result *[][]int) {
	copied := make([]int, len(currentNums))
	copy(copied, currentNums)
	*result = append(*result, copied)

	if len(currentNums) == len(nums) {
		return
	}

	for i := idx; i < len(nums); i++ {
		currentNums = append(currentNums, nums[i])
		backtrack(nums, i+1, currentNums, result)
		currentNums = currentNums[:len(currentNums)-1]
	}
}

func subsetsBacktrack(nums []int) [][]int {
	result := [][]int{}

	backtrack(nums, 0, []int{}, &result)
	return result
}

func TestBacktrack() {
	fmt.Println(subsetsBacktrack([]int{1, 2, 3}))
	fmt.Println(subsetsBacktrack([]int{0}))
}
