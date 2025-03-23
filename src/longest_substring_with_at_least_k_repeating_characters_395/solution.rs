use std::collections::{HashMap, HashSet};

impl Solution {
    pub fn longest_substring(s: String, k: i32) -> i32 {
        let mut char_to_all_occ: HashMap<char, i32> = HashMap::new();

        let char_to_all_occ = s.chars().fold(HashMap::new(), |mut acc, ch| {
            *acc.entry(ch).or_insert(0) += 1;
            acc
        });

        let mut invalid_chars: HashSet<char> = HashSet::new();
        for (&ch, &count) in &char_to_all_occ {
            if count < k {
                invalid_chars.insert(ch);
            }
        }

        if invalid_chars.is_empty() {
            return s.len() as i32;
        }

        let mut window_start = 0;
        let mut result: i32 = 0;

        for (window_end, ch) in s.char_indices() {
            if invalid_chars.contains(&ch) {
                if window_start < window_end {
                    result = result.max(Self::longest_substring(
                        s[window_start..window_end].to_string(),
                        k,
                    ));
                }
                window_start = window_end + 1;
            }
        }

        if window_start < s.len() {
            result = result.max(Self::longest_substring(s[window_start..].to_string(), k));
        }

        result
    }
}
