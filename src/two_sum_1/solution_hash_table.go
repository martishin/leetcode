package two_sum_1

import "fmt"

func twoSumHashTable(nums []int, target int) []int {
	lookupTable := make(map[int]int)

	for idx, num := range nums {
		find := target - num

		if idxFind, exists := lookupTable[find]; exists {
			return []int{idxFind, idx}
		}

		lookupTable[num] = idx
	}

	return []int{-1, 1}
}

func TestHashTable() {
	fmt.Println(twoSumHashTable([]int{3, 2, 4}, 6))
}
