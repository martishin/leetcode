#include <unordered_set>
#include <vector>
#include <unordered_map>
#include <stack>

using namespace std;

class Solution {
public:
    int numberOfComponents(vector<vector<int>>& properties, int k) {
        int n = properties.size();

        unordered_map<int, unordered_set<int>> graph;
        vector<unordered_set<int>> s(n);

        for (int i = 0; i < n; i++) {
            for (auto &x: properties[i]) {
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
            int start = *unvisited.begin();
            dfs(start, graph, unvisited);
            components++;
        }

        return components;
    }

private:
    void dfs(int start, unordered_map<int, unordered_set<int>>& graph, unordered_set<int>& unvisited) {
        stack<int> st;
        st.push(start);
        unvisited.erase(start);

        while (!st.empty()) {
            int node = st.top();
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

    int intersect(const unordered_set<int>& a, const unordered_set<int>& b) {
        int count = 0;
        for (int val : b) {
            if (a.count(val)) {
                count++;
            }
        }

        return count;
    }
};
