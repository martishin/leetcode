package design_tic_tac_toe_348

import (
	"fmt"
)

type TicTacToe struct {
	rows         []int
	columns      []int
	diagonal     int
	antiDiagonal int
	n            int
}

func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows:         make([]int, n),
		columns:      make([]int, n),
		diagonal:     0,
		antiDiagonal: 0,
		n:            n,
	}
}

func (this *TicTacToe) Move(row int, col int, player int) int {
	var sign int
	if player == 1 {
		sign = 1
	} else {
		sign = -1
	}

	this.rows[row] += sign
	this.columns[col] += sign

	if row == col {
		this.diagonal += sign
	}
	if col == this.n-1-row {
		this.antiDiagonal += sign
	}

	if this.rows[row] == sign*this.n || this.columns[col] == sign*this.n || this.diagonal == sign*this.n || this.antiDiagonal == sign*this.n {
		return player
	}

	return 0
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
