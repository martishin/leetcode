from typing import List


class TicTacToe:
    def __init__(self, n: int):
        self.rows: List[int] = [0] * n
        self.cols: List[int] = [0] * n
        self.diagonal: int = 0
        self.anti_diagonal: int = 0

    def move(self, row: int, col: int, player: int) -> int:
        current_player: int = 1 if player == 1 else -1

        self.rows[row] += current_player
        self.cols[col] += current_player

        if row == col:
            self.diagonal += current_player

        if col == (len(self.cols) - row - 1):
            self.anti_diagonal += current_player

        n: int = len(self.rows)

        if (
            abs(self.rows[row]) == n
            or abs(self.cols[col]) == n
            or abs(self.diagonal) == n
            or abs(self.anti_diagonal) == n
        ):
            return player

        return 0


# Your TicTacToe object will be instantiated and called as such:
# obj = TicTacToe(n)
# param_1 = obj.move(row,col,player)
def test():
    obj = TicTacToe(3)
    print(obj.move(0, 0, 1))
    print(obj.move(0, 2, 2))
    print(obj.move(2, 2, 1))
    print(obj.move(1, 1, 2))
    print(obj.move(2, 0, 1))
    print(obj.move(1, 0, 2))
    print(obj.move(2, 1, 1))
