from sortedcontainers import SortedDict


class Solution:
    def __init__(self):
        self._player_to_score = {}
        self._score_to_players_count = SortedDict()

    def add_score(self, player_id: int, score: int):
        old_score = self._player_to_score.get(player_id)

        if old_score:
            self._score_to_players_count[-old_score] -= 1

            new_score = old_score + score
            self._player_to_score[player_id] = new_score
            self._score_to_players_count[-new_score] = (
                self._score_to_players_count.get(-new_score, 0) + 1
            )
        else:
            self._player_to_score[player_id] = score
            self._score_to_players_count[-score] = (
                self._score_to_players_count.get(-score, 0) + 1
            )

    def top(self, k: int) -> int:
        taken, result = 0, 0

        for score, count in self._score_to_players_count.items():
            for _ in range(count):
                if taken == k:
                    break
                result += -score
                taken += 1

        return result

    def reset(self, player_id: int):
        old_score = self._player_to_score[player_id]
        self._score_to_players_count[-old_score] = (
            self._score_to_players_count.get(-old_score, 0) - 1
        )
        del self._player_to_score[player_id]


def test():
    solution = Solution()
    solution.add_score(1, 73)
    solution.add_score(2, 56)
    solution.add_score(3, 39)
    solution.add_score(4, 51)
    solution.add_score(5, 4)
    print(solution.top(1))
    solution.reset(1)
    solution.reset(2)
    solution.add_score(2, 51)
    print(solution.top(3))
