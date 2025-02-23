package separate_squares_i_3453

import "sort"

func separateSquares(squares [][]int) float64 {
	sort.Slice(squares, func(i, j int) bool {
		if squares[i][1] == squares[j][1] {
			return squares[i][0] < squares[j][0]
		}
		return squares[i][1] < squares[j][1]
	})

	totalSum := sumOfSquares(squares)

	l := 0.0
	r := 1e9

	for (r - l) > 1e-5 {
		lowerSum := 0.0
		mid := (l + r) / 2.0
		for _, square := range squares {
			y, l := float64(square[1]), float64(square[2])
			if y > mid {
				break
			} else if y+l > mid {
				lowerSum += (mid - y) * l
			} else {
				lowerSum += l * l
			}
		}
		upperSum := totalSum - lowerSum
		if lowerSum >= upperSum {
			r = mid
		} else {
			l = mid
		}
	}
	return r
}

func sumOfSquares(squares [][]int) float64 {
	sum := 0.0
	for _, square := range squares {
		sum += float64(square[2] * square[2])
	}

	return sum
}
