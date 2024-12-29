pub struct Solution;

impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        let n = strs.len();
        if n == 0 {
            return String::new();
        }

        let mut pointer = 0;
        let mut result_len = 0;

        loop {
            if pointer >= strs[0].len() {
                break;
            }

            let prefix_char = strs[0].chars().nth(pointer).unwrap();

            for i in 0..n {
                if pointer >= strs[i].len() || strs[i].chars().nth(pointer).unwrap() != prefix_char
                {
                    return strs[0][0..result_len].to_string();
                }
            }

            pointer += 1;
            result_len += 1;
        }

        strs[0][0..result_len].to_string()
    }
}

pub fn test() {
    println!(
        "{}",
        Solution::longest_common_prefix(
            vec!["flower", "flow", "flight"]
                .into_iter()
                .map(String::from)
                .collect()
        )
    );
    println!(
        "{}",
        Solution::longest_common_prefix(
            vec!["dog", "racecar", "car"]
                .into_iter()
                .map(String::from)
                .collect()
        )
    );
}
