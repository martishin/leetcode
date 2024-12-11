package transpose_matrix_867

func transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)
	}

	for r, row := range matrix {
		for c := range row {
			result[c][r] = matrix[r][c]
		}
	}

	return result
}
