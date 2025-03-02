package longest_subarray_of_1_s_after_deleting_one_element_1493

func longestSubarray(nums []int) int {
	count := 0
	result := 0
	l := 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			count++
		}

		for count > 1 {
			if nums[l] == 0 {
				count--
			}
			l++
		}

		result = max(result, r-l)
	}

	return result
}
