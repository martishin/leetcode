class TicTacToe {
  private rows: number[];
  private cols: number[];
  private diagonal: number;
  private antiDiagonal: number;

  constructor(n: number) {
    this.rows = new Array(n).fill(0);
    this.cols = new Array(n).fill(0);
    this.diagonal = 0;
    this.antiDiagonal = 0;
  }

  move(row: number, col: number, player: number): number {
    const currentPlayer = player === 1 ? 1 : -1;

    this.rows[row] += currentPlayer;
    this.cols[col] += currentPlayer;

    if (row === col) {
      this.diagonal += currentPlayer;
    }

    if (col === this.cols.length - row - 1) {
      this.antiDiagonal += currentPlayer;
    }

    const n = this.rows.length;

    if (
      Math.abs(this.rows[row]) === n ||
      Math.abs(this.cols[col]) === n ||
      Math.abs(this.diagonal) === n ||
      Math.abs(this.antiDiagonal) === n
    ) {
      return player;
    }

    return 0;
  }
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * var obj = new TicTacToe(n);
 * var param_1 = obj.move(row, col, player);
 */
export function test(): void {
  const game = new TicTacToe(3);

  console.log(game.move(0, 0, 1));
  console.log(game.move(0, 2, 2));
  console.log(game.move(2, 2, 1));
  console.log(game.move(1, 1, 2));
  console.log(game.move(2, 0, 1));
  console.log(game.move(2, 1, 1));
}
