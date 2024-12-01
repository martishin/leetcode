from collections import Counter
from typing import List


class Solution:
    def longestPalindrome(self, words: List[str]) -> int:
        word_to_count = Counter(words)

        visited = set()
        used_center = False
        result = 0

        for word, count in word_to_count.items():
            if word in visited:
                continue

            visited.add(word)

            if word[0] == word[1]:
                if count % 2 == 0:
                    result += count * 2
                else:
                    result += (count - 1) * 2
                    if not used_center:
                        used_center = True
                        result += 2
            else:
                pair = word[1] + word[0]
                pair_count = word_to_count[pair]

                if pair_count > 0:
                    visited.add(pair)
                    result += 2 * 2 * min(count, pair_count)

        return result


def test():
    solution = Solution()
    print(solution.longestPalindrome(["lc", "cl", "gg"]))
    print(solution.longestPalindrome(["ab", "ty", "yt", "lc", "cl", "ab"]))
    print(solution.longestPalindrome(["cc", "ll", "xx"]))
