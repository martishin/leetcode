use std::cmp::max;

pub struct SolutionDp;

impl SolutionDp {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let n = nums.len();

        if n == 0 {
            return 0;
        } else if n == 1 {
            return nums[0];
        }

        let mut dp = vec![0; n + 1];

        dp[0] = nums[0];
        dp[1] = max(nums[0], nums[1]);

        for i in 2..n {
            dp[i] = max(dp[i - 1], dp[i - 2] + nums[i]);
        }

        dp[n - 1]
    }
}

pub fn test() {
    println!("{}", SolutionDp::rob(vec![1, 2, 3, 1]));
    println!("{}", SolutionDp::rob(vec![2, 7, 9, 3, 1]));
    println!("{}", SolutionDp::rob(vec![]));
}
