pub struct Solution;

impl Solution {
    // Convert an IP string to a 32-bit integer
    fn ip_to_int(ip: &str) -> u32 {
        ip.split('.')
            .map(|part| part.parse::<u32>().unwrap())
            .fold(0, |acc, octet| (acc << 8) + octet)
    }

    // Convert a 32-bit integer to an IP string
    fn int_to_ip(num: u32) -> String {
        (0..=3)
            .rev()
            .map(|i| ((num >> (i * 8)) & 0xFF).to_string())
            .collect::<Vec<_>>()
            .join(".")
    }

    // Generate CIDR notation from IP and prefix length
    fn to_cidr(ip: u32, bits: u32) -> String {
        format!("{}/{}", Self::int_to_ip(ip), bits)
    }

    // Main function to calculate the largest block and generate CIDR
    pub fn ip_to_cidr(ip: String, mut n: i32) -> Vec<String> {
        let mut cidrs = Vec::new();
        let mut num = Self::ip_to_int(&ip);

        while n > 0 {
            // Calculate the largest possible block size
            let zeros = num.trailing_zeros().min(31 - (n as u32).leading_zeros());

            // Determine the block size and generate the CIDR
            let cnt = 1u32 << zeros;
            cidrs.push(Self::to_cidr(num, 32 - zeros));

            // Move to the next block
            num += cnt;
            n -= cnt as i32;
        }

        cidrs
    }
}

pub fn test() {
    println!("{:?}", Solution::ip_to_cidr("255.0.0.7".to_string(), 10));
    println!(
        "{:?}",
        Solution::ip_to_cidr("117.145.102.62".to_string(), 8)
    );
    println!("{:?}", Solution::ip_to_cidr("0.0.0.0".to_string(), 1));
}
