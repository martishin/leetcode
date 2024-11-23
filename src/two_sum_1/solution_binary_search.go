package two_sum_1

import (
	"fmt"
	"sort"
)

type Pair struct {
	num int
	idx int
}

func twoSumBinarySearch(nums []int, target int) []int {
	numsWithIndices := make([]Pair, len(nums))
	for i, num := range nums {
		numsWithIndices[i] = Pair{num: num, idx: i}
	}

	sort.Slice(numsWithIndices, func(i, j int) bool {
		return numsWithIndices[i].num < numsWithIndices[j].num
	})

	for idx, pair := range numsWithIndices {
		num := pair.num
		originalIdx := pair.idx
		find := target - num

		l := idx + 1
		r := len(numsWithIndices)
		for l < r {
			m := (l + r) / 2
			if numsWithIndices[m].num == find {
				return []int{originalIdx, numsWithIndices[m].idx}
			} else if numsWithIndices[m].num < find {
				l = m + 1
			} else {
				r = m
			}
		}
	}

	return []int{-1, 1}
}

func TestBinarySearch() {
	fmt.Println(twoSumBinarySearch([]int{3, 2, 4}, 6))
}
