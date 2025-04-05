#ifndef SOLUTION_H
#define SOLUTION_H

#include <unordered_map>
#include <unordered_set>
#include <vector>

class Solution {
public:
    static int numberOfComponents(std::vector<std::vector<int>>& properties, int k);

private:
    static void dfs(int start, std::unordered_map<int, std::unordered_set<int>>& graph,
                    std::unordered_set<int>& unvisited);
    static int intersect(const std::unordered_set<int>& a, const std::unordered_set<int>& b);
};

void test();

#endif  // SOLUTION_H
