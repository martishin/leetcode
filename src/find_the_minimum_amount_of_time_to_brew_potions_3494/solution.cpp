#include <vector>
#include <iostream>

using namespace std;

class Solution {
public:
    long long minTime(vector<int>& skill, vector<int>& mana) {
        auto m = mana.size();
        auto n = skill.size();

        vector<long long> dp(n + 1, 0);

        for (auto i = 0; i < m; i++) {
            for (auto j = 0; j < n; j++) {
                dp[j + 1] = max(dp[j + 1], dp[j]) + mana[i] * skill[j];
            }
            for (auto j = n - 1; j > 0; j--) {
                dp[j] = dp[j + 1] - mana[i] * skill[j];
            }
        }

        return dp[n];
    }
};
