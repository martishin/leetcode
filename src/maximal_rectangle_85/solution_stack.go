package maximal_rectangle_85

import "fmt"

func maximalRectangleStack(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])

	if rows == 0 || cols == 0 {
		return 0
	}

	maxRectangle := 0
	heights := make([]int, cols)
	for row := range rows {
		for col := range cols {
			if matrix[row][col] == '1' {
				heights[col]++
			} else {
				heights[col] = 0
			}
		}
		maxRectangle = max(maxRectangle, maximalHistogram(heights))
	}
	return maxRectangle
}

func maximalHistogram(heights []int) int {
	maxHistogram := 0
	stack := []int{-1}

	for r := range heights {
		for len(stack) > 1 && heights[stack[len(stack)-1]] > heights[r] {
			maxHistogram = max(maxHistogram, heights[stack[len(stack)-1]]*(r-stack[len(stack)-2]-1))
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, r)
	}

	for len(stack) > 1 {
		maxHistogram = max(maxHistogram, heights[stack[len(stack)-1]]*(len(heights)-stack[len(stack)-2]-1))
		stack = stack[:len(stack)-1]
	}

	return maxHistogram
}

func TestStack() {
	fmt.Println(maximalRectangleStack([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))

	fmt.Println(maximalRectangleStack([][]byte{
		{'0'},
	}))

	fmt.Println(maximalRectangleStack([][]byte{
		{'1'},
	}))
}
