package house_robber_198

import "fmt"

func robDp(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], nums[i]+dp[i-2])
	}

	return dp[n-1]
}

func TestDp() {
	fmt.Println(robDp([]int{1, 2, 3, 1}))
	fmt.Println(robDp([]int{2, 7, 9, 3, 1}))
}
