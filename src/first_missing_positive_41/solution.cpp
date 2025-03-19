class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        auto n = nums.size();

        for (auto i = 0; i < nums.size(); i++) {
            while (nums[i] <= nums.size() && nums[i] > 0 && nums[i] != nums[nums[i] - 1]) {
                std::swap(nums[i], nums[nums[i] - 1]);
            }
        }

        for (auto v : nums) {
            cout << v << endl;
        }

        for (auto i = 0; i < n; i++) {
            if (nums[i] != i + 1) {
                return i + 1;
            }
        }

        return n + 1;
    }
};
