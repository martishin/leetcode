#include "solution.h"

#include <iostream>
#include <stack>
#include <unordered_map>
#include <unordered_set>
#include <vector>

using namespace std;

int Solution::numberOfComponents(vector<vector<int>>& properties, int k) {
    const auto n = properties.size();

    unordered_map<int, unordered_set<int>> graph;
    vector<unordered_set<int>> s(n);

    for (int i = 0; i < n; i++) {
        for (auto& x : properties[i]) {
            s[i].insert(x);
        }
    }

    for (int i = 0; i < n; ++i) {
        for (int j = i + 1; j < n; ++j) {
            if (intersect(s[i], s[j]) >= k) {
                graph[i].insert(j);
                graph[j].insert(i);
            }
        }
    }

    unordered_set<int> unvisited;
    for (int i = 0; i < n; ++i) {
        unvisited.insert(i);
    }

    int components = 0;
    while (!unvisited.empty()) {
        const auto start = *unvisited.begin();
        dfs(start, graph, unvisited);
        components++;
    }

    return components;
}

void Solution::dfs(int start, unordered_map<int, unordered_set<int>>& graph, unordered_set<int>& unvisited) {
    stack<int> st;
    st.push(start);
    unvisited.erase(start);

    while (!st.empty()) {
        const auto node = st.top();
        st.pop();

        for (auto nb : graph[node]) {
            auto it = unvisited.find(nb);
            if (it != unvisited.end()) {
                unvisited.erase(it);
                st.push(nb);
            }
        }
    }
}

int Solution::intersect(const unordered_set<int>& a, const unordered_set<int>& b) {
    int count = 0;
    for (const auto val : b) {
        if (a.count(val)) {
            ++count;
        }
    }

    return count;
}

void test() {
    vector<vector<int>> input = {{1, 2}, {1, 1}, {3, 4}, {4, 5}, {5, 6}, {7, 7}};

    const auto result = Solution::numberOfComponents(input, 1);
    cout << result << endl;
}
