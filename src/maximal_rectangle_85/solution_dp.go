package maximal_rectangle_85

import "fmt"

func maximalRectangleDP(matrix [][]byte) int {
	rows := len(matrix)

	if rows == 0 {
		return 0
	}

	cols := len(matrix[0])

	height := make([]int, cols)
	left := make([]int, cols)
	right := make([]int, cols)

	for i := range right {
		right[i] = cols
	}

	maxRectangle := 0
	for row := range rows {
		currLeft := 0
		currRight := cols

		for col := range cols {
			if matrix[row][col] == '1' {
				height[col]++
			} else {
				height[col] = 0
			}
		}

		for col := range cols {
			if matrix[row][col] == '1' {
				left[col] = max(left[col], currLeft)
			} else {
				left[col] = 0
				currLeft = col + 1
			}
		}

		for col := cols - 1; col >= 0; col-- {
			if matrix[row][col] == '1' {
				right[col] = min(right[col], currRight)
			} else {
				right[col] = cols
				currRight = col
			}
		}

		for col := range cols {
			maxRectangle = max(maxRectangle, height[col]*(right[col]-left[col]))
		}
	}

	return maxRectangle
}

func TestDP() {
	fmt.Println(maximalRectangleDP([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))

	fmt.Println(maximalRectangleDP([][]byte{
		{'0'},
	}))

	fmt.Println(maximalRectangleDP([][]byte{
		{'1'},
	}))
}
