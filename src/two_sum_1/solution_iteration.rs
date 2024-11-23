pub struct Solution;

struct Pair {
    num: i32,
    idx: usize,
}

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut nums_with_indices: Vec<Pair> = nums
            .iter()
            .enumerate()
            .map(|(idx, &num)| Pair { num, idx })
            .collect();

        nums_with_indices.sort_by_key(|pair| pair.num);

        let mut l = 0;
        let mut r = nums_with_indices.len() - 1;

        while l < r {
            let sum = nums_with_indices[l].num + nums_with_indices[r].num;

            match sum.cmp(&target) {
                std::cmp::Ordering::Less => l += 1,
                std::cmp::Ordering::Greater => r -= 1,
                std::cmp::Ordering::Equal => {
                    return vec![
                        nums_with_indices[l].idx as i32,
                        nums_with_indices[r].idx as i32,
                    ];
                }
            }
        }

        vec![-1, -1]
    }
}

pub fn test() {
    println!("{:?}", Solution::two_sum(vec![3, 2, 4], 6));
}
