package rotate_image_48

import "fmt"

func rotateGroups(matrix [][]int) {
	n := len(matrix)
	shift := 0
	for i := n; i > 1; i -= 2 {
		for j := 0; j < i-1; j++ {
			matrix[shift][shift+j], matrix[shift+j][n-1-shift] = matrix[shift+j][n-1-shift], matrix[shift][shift+j]
			matrix[shift][shift+j], matrix[n-1-shift][n-1-shift-j] = matrix[n-1-shift][n-1-shift-j], matrix[shift][shift+j]
			matrix[shift][shift+j], matrix[n-1-shift-j][shift] = matrix[n-1-shift-j][shift], matrix[shift][shift+j]
		}
		shift += 1
	}
}

func rotateTranspose(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		l, r := 0, n-1
		for l < r {
			matrix[i][l], matrix[i][r] = matrix[i][r], matrix[i][l]
			l++
			r--
		}
	}
}

func Test() {
	test1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateTranspose(test1)
	fmt.Println(test1)

	test2 := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotateTranspose(test2)
	fmt.Println(test2)
}
