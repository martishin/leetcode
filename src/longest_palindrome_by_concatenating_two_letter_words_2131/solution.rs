use std::collections::{HashMap, HashSet};

pub struct Solution;

impl Solution {
    pub fn longest_palindrome(words: Vec<String>) -> i32 {
        let mut word_to_count = HashMap::new();

        for word in &words {
            *word_to_count.entry(word.as_str()).or_insert(0) += 1;
        }

        let mut visited: HashSet<String> = HashSet::new();
        let mut used_center = false;
        let mut result = 0;

        for (&word, &count) in &word_to_count {
            if !visited.insert(word.to_string()) {
                continue;
            }

            let chars: Vec<char> = word.chars().collect();
            if chars[0] == chars[1] {
                if count % 2 == 0 {
                    result += count * 2;
                } else {
                    result += (count - 1) * 2;
                    if !used_center {
                        used_center = true;
                        result += 2;
                    }
                }
            } else {
                let pair = format!("{}{}", chars[1], chars[0]);
                if let Some(&pair_count) = word_to_count.get(pair.as_str()) {
                    visited.insert(pair);
                    result += 2 * 2 * count.min(pair_count);
                }
            }
        }

        result
    }
}

pub fn test() {
    println!(
        "{}",
        Solution::longest_palindrome(vec!["lc".to_string(), "cl".to_string(), "gg".to_string()])
    );
    println!(
        "{}",
        Solution::longest_palindrome(vec![
            "ab".to_string(),
            "ty".to_string(),
            "yt".to_string(),
            "lc".to_string(),
            "cl".to_string(),
            "ab".to_string()
        ])
    );
    println!(
        "{}",
        Solution::longest_palindrome(vec!["cc".to_string(), "ll".to_string(), "xx".to_string()])
    );
}
