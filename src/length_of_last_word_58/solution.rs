pub struct Solution;

impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        let mut result = 0;
        for i in (0..s.len()).rev() {
            if &s[i..i + 1] != " " {
                result += 1;
            } else if result != 0 {
                break;
            }
        }
        result
    }
}

pub fn test() {
    println!(
        "{}",
        Solution::length_of_last_word("Hello World".to_string())
    );
    println!(
        "{}",
        Solution::length_of_last_word("   fly me   to   the moon  ".to_string())
    );
    println!(
        "{}",
        Solution::length_of_last_word("luffy is still joyboy".to_string())
    );
}
