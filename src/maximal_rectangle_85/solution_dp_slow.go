package maximal_rectangle_85

import "math"

func maximalRectangleDPSlow(matrix [][]byte) int {
	n := len(matrix)
	m := len(matrix[0])

	if n == 0 || m == 0 {
		return 0
	}

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, m)
	}

	result := math.MinInt32

	for i := range n {
		for j := range m {
			top := 0
			left := 0
			topLeft := 0

			if i > 0 && j > 0 {
				top = dp[i-1][j]
				left = dp[i][j-1]
				topLeft = dp[i-1][j-1]
			} else if i > 0 {
				top = dp[i-1][j]
			} else if j > 0 {
				left = dp[i][j-1]
			}

			dp[i][j] = int(matrix[i][j]-'0') + top + left - topLeft
		}
	}

	for i1 := 0; i1 < n; i1++ {
		for j1 := 0; j1 < m; j1++ {
			for i2 := i1; i2 < n; i2++ {
				for j2 := j1; j2 < m; j2++ {
					left := 0
					top := 0
					topLeft := 0

					if i1 > 0 && j1 > 0 {
						top = dp[i1-1][j2]
						left = dp[i2][j1-1]
						topLeft = dp[i1-1][j1-1]
					} else if i1 > 0 {
						top = dp[i1-1][j2]
					} else if j1 > 0 {
						left = dp[i2][j1-1]
					}

					size := dp[i2][j2] - left - top + topLeft
					if size == (j2-j1+1)*(i2-i1+1) {
						result = max(result, size)
					}
				}
			}
		}
	}

	if result == math.MinInt32 {
		return 0
	}

	return result
}
