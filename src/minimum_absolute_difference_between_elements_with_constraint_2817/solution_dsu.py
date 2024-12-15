from collections import Counter
from typing import List


# def min_diff(nums: List[int], x: int) -> int:
#     sorted_list = sorted((val, idx) for idx, val in enumerate(nums))
#     n = len(nums)
#
#     parent = list(range(n))  # DSU parent array
#     size = [1] * n  # DSU size array
#     active = [False] * n  # Active elements tracker
#
#     def find(node):
#         if parent[node] != node:
#             parent[node] = find(parent[node])  # Path compression
#         return parent[node]
#
#     def union(node1, node2):
#         root1 = find(node1)
#         root2 = find(node2)
#
#         if root1 != root2:
#             # Union bt size
#             if size[root1] < size[root2]:
#                 parent[root1] = root2
#                 size[root2] += size[root1]
#             else:
#                 parent[root2] = root1
#                 size[root1] += size[root2]
#
#     result = 1e10
#
#     for val, idx in sorted_list:
#         # Activate the current element
#         active[idx] = True
#
#         # Union with neighbors if they are active
#         if idx > 0 and active[idx - 1]:
#             union(idx, idx - 1)
#         if idx < n - 1 and active[idx + 1]:
#             union(idx, idx + 1)
#
#         # Find closest neighbors
#         root = find(idx)
#         if root > x - 1:
#             result = min(result, abs(nums[idx] - nums[root - x]))
#         if root < n - x:
#             result = min(result, abs(nums[idx] - nums[root + x]))
#
#     return result


def min_diff(nums: List[int], x: int) -> int:
    # Add sentinel values to handle boundaries.
    # This helps ensure that when we move in direction d,
    # we do not go out of range.
    nums += [float("inf"), float("-inf")]

    # Create a sorted list of unique values from nums.
    # 'rnk' maps each value to an index (rank).
    rnk = sorted(set(nums))

    # We will perform two passes:
    # One with d = 1 (moving forward in the rank array)
    # and one with d = -1 (moving backward in the rank array).
    # This ensures we find the minimal difference in both directions.

    res = float("inf")
    n = len(nums)

    for d in [1, -1]:
        # dp will map each value to its current index in 'rnk'.
        dp = {val: i for i, val in enumerate(rnk)}

        # 'active' tracks how many occurrences of each value are currently considered "active."
        # Initially, we consider all elements from nums[x:] as active.
        # As we iterate, we'll slide this window by deactivating one element at a time.
        active = Counter(nums[x:])

        def find(num):
            """
            The find function attempts to find an active rank index for 'num'.
            We look up dp[num] to find its current rank index.
            If the value at that rank is not active, we move one step in direction d
            and recursively find a new dp[num].
            This mimics a path-compression-like behavior specialized for this problem.
            """
            current_rank_index = dp[num]
            next_val = rnk[current_rank_index]

            # If the current value at this rank isn't active,
            # try moving one step in direction d and update dp accordingly.
            if not active[next_val]:
                dp[num] = find(rnk[current_rank_index + d])
            return dp[num]

        # We iterate up to n - x - 2 because we are considering pairs at least x apart.
        # On each iteration:
        # 1. We find the closest active rank for nums[i].
        # 2. Update the result with the minimal absolute difference.
        # 3. Deactivate nums[i + x] to slide the window.
        for i in range(n - x - 2):
            closest_rank_index = find(nums[i])
            closest_val = rnk[closest_rank_index]
            res = min(res, abs(closest_val - nums[i]))

            # Deactivate one occurrence of nums[i + x]
            active[nums[i + x]] -= 1

    return res


def test():
    print(min_diff([4, 3, 2, 4], 2))
    print(min_diff([5, 3, 2, 10, 15], 1))
    print(min_diff([1, 2, 3, 4], 3))
    print(min_diff([14, 111, 16], 1))
