package permutations_47

import "fmt"

func backtrack(nums []int, currentNums []int, permutations *[][]int) {
	if len(currentNums) == len(nums) {
		copied := make([]int, len(nums))
		copy(copied, currentNums)
		*permutations = append(*permutations, copied)
		return
	}

	for _, num := range nums {
		used := false

		for _, currentNum := range currentNums {
			if num == currentNum {
				used = true
				break
			}
		}

		if !used {
			currentNums = append(currentNums, num)
			backtrack(nums, currentNums, permutations)
			currentNums = currentNums[:len(currentNums)-1]
		}
	}
}

func permute(nums []int) [][]int {
	permutations := [][]int{}

	backtrack(nums, []int{}, &permutations)

	return permutations
}

func Test() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
	fmt.Println(permute([]int{1}))
}
