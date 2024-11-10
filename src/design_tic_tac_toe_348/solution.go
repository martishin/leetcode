package design_tic_tac_toe_348

import (
	"fmt"
	"math"
)

type TicTacToe struct {
	rows         []int
	cols         []int
	diagonal     int
	antiDiagonal int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows:         make([]int, n),
		cols:         make([]int, n),
		diagonal:     0,
		antiDiagonal: 0,
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	currentPlayer := 1
	if player != 1 {
		currentPlayer = -1
	}

	this.rows[row] += currentPlayer
	this.cols[col] += currentPlayer

	if row == col {
		this.diagonal += currentPlayer
	}

	if col == len(this.cols)-row-1 {
		this.antiDiagonal += currentPlayer
	}

	n := len(this.rows)

	if abs(this.rows[row]) == n || abs(this.cols[col]) == n || abs(this.diagonal) == n || abs(this.antiDiagonal) == n {
		return player
	}

	return 0
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

// Test
/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
func Test() {
	obj := Constructor(3)
	fmt.Println(obj.Move(0, 0, 1))
	fmt.Println(obj.Move(0, 2, 2))
	fmt.Println(obj.Move(2, 2, 1))
	fmt.Println(obj.Move(1, 1, 2))
	fmt.Println(obj.Move(2, 0, 1))
	fmt.Println(obj.Move(1, 0, 2))
	fmt.Println(obj.Move(2, 1, 1))
}
