pub struct Solution;

impl Solution {
    pub fn candy(ratings: Vec<i32>) -> i32 {
        let n = ratings.len();

        let mut left_to_right: Vec<i32> = vec![1; n];

        for i in 1..n {
            if ratings[i] > ratings[i - 1] {
                left_to_right[i] = left_to_right[i - 1] + 1;
            }
        }

        for i in (0..n - 1).rev() {
            if ratings[i] > ratings[i + 1] {
                left_to_right[i] = std::cmp::max(left_to_right[i], left_to_right[i + 1] + 1);
            }
        }

        left_to_right.iter().sum()
    }
}

pub fn test() {
    println!("{}", Solution::candy(vec![1, 0, 2]));
    println!("{}", Solution::candy(vec![1, 2, 2]));
}
