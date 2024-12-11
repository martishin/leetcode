from typing import List


class Solution:
    def rotateTheBox(self, boxGrid: List[List[str]]) -> List[List[str]]:
        m = len(boxGrid)
        n = len(boxGrid[0])

        for row in range(m):
            right_border = n - 1
            for col in range(n - 1, -1, -1):
                obj = boxGrid[row][col]
                if obj == "*":
                    right_border = col - 1
                elif boxGrid[row][col] == "#":
                    if col == right_border:
                        right_border -= 1
                        continue
                    boxGrid[row][col] = "."
                    boxGrid[row][right_border] = "#"
                    right_border -= 1

        result = [["." for _ in range(m)] for _ in range(n)]

        for row in range(m):
            for col in range(n):
                result[col][m - 1 - row] = boxGrid[row][col]

        return result


def test():
    solution = Solution()
    print(solution.rotateTheBox([["#", ".", "#"]]))
    print(solution.rotateTheBox([["#", ".", "*", "."], ["#", "#", "*", "."]]))
    print(
        solution.rotateTheBox(
            [
                ["#", "#", "*", ".", "*", "."],
                ["#", "#", "#", "*", ".", "."],
                ["#", "#", "#", ".", "#", "."],
            ]
        )
    )
