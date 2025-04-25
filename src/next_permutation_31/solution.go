package next_permutation_31

import "fmt"

func nextPermutation(nums []int) {
	i := len(nums) - 2

	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := len(nums) - 1
		for nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1)
}

func reverse(nums []int, l int) {
	r := len(nums) - 1

	for l < r {
		swap(nums, l, r)
		l++
		r--
	}
}

func swap(nums []int, i, j int) {
	nums[i] += nums[j]
	nums[j] = nums[i] - nums[j]
	nums[i] -= nums[j]
}

func Test() {
	t1 := []int{1, 2, 3}
	nextPermutation(t1)
	fmt.Println(t1)

	t2 := []int{3, 2, 1}
	nextPermutation(t2)
	fmt.Println(t2)

	t3 := []int{1, 1, 5}
	nextPermutation(t3)
	fmt.Println(t3)
}
