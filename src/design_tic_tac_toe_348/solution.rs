struct TicTacToe {
    rows: Vec<i32>,
    cols: Vec<i32>,
    diagonal: i32,
    anti_diagonal: i32,
}

impl TicTacToe {
    fn new(n: i32) -> Self {
        let size = n as usize;

        TicTacToe {
            rows: vec![0; size],
            cols: vec![0; size],
            diagonal: 0,
            anti_diagonal: 0,
        }
    }

    fn make_a_move(&mut self, row: i32, col: i32, player: i32) -> i32 {
        let current_player = if player == 1 { 1 } else { -1 };

        let row_index = row as usize;
        let col_index = col as usize;

        self.rows[row_index] += current_player;
        self.cols[col_index] += current_player;

        if row_index == col_index {
            self.diagonal += current_player;
        }

        if col_index == (self.cols.len() - row_index - 1) {
            self.anti_diagonal += current_player;
        }

        let n = self.rows.len() as i32;

        if self.rows[row_index].abs() == n
            || self.cols[col_index].abs() == n
            || self.diagonal.abs() == n
            || self.anti_diagonal.abs() == n
        {
            return player;
        }

        0
    }
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * let obj = TicTacToe::new(n);
 * let ret_1: i32 = obj.move(row, col, player);
 */
pub fn test() {
    let mut obj = TicTacToe::new(3);
    println!("{}", obj.make_a_move(0, 0, 1));
    println!("{}", obj.make_a_move(0, 2, 2));
    println!("{}", obj.make_a_move(2, 2, 1));
    println!("{}", obj.make_a_move(1, 1, 2));
    println!("{}", obj.make_a_move(2, 0, 1));
    println!("{}", obj.make_a_move(1, 0, 2));
    println!("{}", obj.make_a_move(2, 1, 1));
}
