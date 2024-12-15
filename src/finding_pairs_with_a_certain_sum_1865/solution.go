package finding_pairs_with_a_certain_sum_1865

import (
	"fmt"
	"sort"
)

type FindSumPairs struct {
	nums1        []int
	nums2        []int
	nums2ToCount map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	nums2ToCount := make(map[int]int)

	for _, num2 := range nums2 {
		nums2ToCount[num2]++
	}

	return FindSumPairs{
		nums1:        nums1,
		nums2:        nums2,
		nums2ToCount: nums2ToCount,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	this.nums2ToCount[this.nums2[index]]--
	this.nums2[index] += val
	this.nums2ToCount[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
	result := 0
	for _, num1 := range this.nums1 {
		if num1 >= tot {
			break
		}

		result += this.nums2ToCount[tot-num1]
	}

	return result
}

// Test
/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
func Test() {
	obj := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	fmt.Println(obj.Count(7))
	obj.Add(3, 2)
	fmt.Println(obj.Count(8))
	fmt.Println(obj.Count(4))
	obj.Add(0, 1)
	obj.Add(1, 1)
	fmt.Println(obj.Count(7))
}
