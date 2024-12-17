package rotating_the_box_1861

import "fmt"

func rotateTheBox(boxGrid [][]byte) [][]byte {
	n := len(boxGrid)
	m := len(boxGrid[0])

	resultGrid := make([][]byte, m)
	for i := range resultGrid {
		resultGrid[i] = make([]byte, n)
		for j := range resultGrid[i] {
			resultGrid[i][j] = '.'
		}
	}

	for i, row := range boxGrid {
		last := m - 1
		for j := len(row) - 1; j >= 0; j-- {
			ch := row[j]
			switch ch {
			case '#':
				resultGrid[last][n-1-i] = '#'
				last--
			case '*':
				resultGrid[j][n-1-i] = '*'
				last = j - 1
			}
		}
	}

	return resultGrid
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Printf("%c ", col)
		}
		fmt.Println()
	}
	fmt.Println()
}

func Test() {
	grid1 := [][]byte{
		{'#', '.', '#'},
	}
	printGrid(rotateTheBox(grid1))

	grid2 := [][]byte{
		{'#', '.', '*', '.'},
		{'#', '#', '*', '.'},
	}
	printGrid(rotateTheBox(grid2))

	grid3 := [][]byte{
		{'#', '#', '*', '.', '*', '.'},
		{'#', '#', '#', '*', '.', '.'},
		{'#', '#', '#', '.', '#', '.'},
	}
	printGrid(rotateTheBox(grid3))
}
