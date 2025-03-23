impl Solution {
    pub fn number_of_special_substrings(s: String) -> i32 {
        let mut result: i32 = 0;
        let mut start: usize = 0;
        let mut freq = [0; 26];

        let bytes = s.as_bytes();

        for (end, &ch) in bytes.iter().enumerate() {
            let idx = (ch - b'a') as usize;
            freq[idx] += 1;

            while freq[idx] > 1 {
                let start_idx = (bytes[start] - b'a') as usize;
                freq[start_idx] -= 1;
                start += 1;
            }

            result += (end - start + 1) as i32;
        }

        result
    }
}
