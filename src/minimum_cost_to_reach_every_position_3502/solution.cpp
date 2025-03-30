using namespace std;

class Solution {
public:
    vector<int> minCosts(vector<int>& cost) {
        vector<int> result;
        int min_value = numeric_limits<int>::max();

        for (auto c : cost) {
            min_value = min(min_value, c);
            result.push_back(min_value);
        }

        return result;
    }
};
