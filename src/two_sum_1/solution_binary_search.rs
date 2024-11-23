pub struct Solution;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut nums_with_indices: Vec<(i32, usize)> = nums
            .iter()
            .enumerate()
            .map(|(idx, &num)| (num, idx))
            .collect();

        nums_with_indices.sort_by_key(|&(num, _)| num);

        for (idx, &(num, original_idx)) in nums_with_indices.iter().enumerate() {
            let mut l = idx + 1;
            let mut r = nums_with_indices.len();

            let find = target - num;
            while l < r {
                let m = (l + r) / 2;
                match nums_with_indices[m].0.cmp(&find) {
                    std::cmp::Ordering::Equal => {
                        return vec![original_idx as i32, nums_with_indices[m].1 as i32];
                    }
                    std::cmp::Ordering::Less => {
                        l = m + 1;
                    }
                    std::cmp::Ordering::Greater => {
                        r = m;
                    }
                }
            }
        }

        vec![-1, -1]
    }
}

pub fn test() {
    println!("{:?}", Solution::two_sum(vec![3, 2, 4], 6));
}
