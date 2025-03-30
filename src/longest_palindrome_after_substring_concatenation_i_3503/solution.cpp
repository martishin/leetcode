using namespace std;

class Solution {
public:
    int longestPalindrome(string s, string t) {
        int max_len = 0;

        for (int s_idx = 0; s_idx <= s.size(); ++s_idx) {
            for (int t_idx = 0; t_idx <= t.size(); ++t_idx) {
                int virtual_len = s_idx + (t.size() - t_idx);

                for (int center = 0; center < virtual_len; ++center) {
                    int len1 = expandAroundCenter(s, t, s_idx, t_idx, center, center);
                    int len2 = expandAroundCenter(s, t, s_idx, t_idx, center, center + 1);
                    max_len = max({max_len, len1, len2});
                }
            }
        }

        return max_len;
    }

private:
    char getChar(const string& s, const string& t, int s_idx, int t_idx, int index) {
        if (index < s_idx) return s[index];
        return t[t_idx + (index - s_idx)];
    }

    int expandAroundCenter(const string& s, const string& t, int s_idx, int t_idx, int left, int right) {
        int max_virtual_len = s_idx + (int)t.size() - t_idx;
        while (left >= 0 && right < max_virtual_len) {
            char left_char = getChar(s, t, s_idx, t_idx, left);
            char right_char = getChar(s, t, s_idx, t_idx, right);
            if (left_char != right_char) {
                break;
            }
            --left;
            ++right;
        }
        return right - left - 1;
    }
};
