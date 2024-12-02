pub struct Solution;

impl Solution {
    fn ip_to_int(ip: &str) -> u32 {
        let mut result = 0u32;
        for part in ip.split('.') {
            let octet: u32 = part.parse().unwrap();
            result = (result << 8) + octet;
        }
        result
    }

    fn int_to_ip(num: u32) -> String {
        let mut octets = Vec::new();
        for i in (0..=3).rev() {
            let octet = (num >> (i * 8)) & 0xFF;
            octets.push(octet.to_string());
        }
        octets.join(".")
    }

    fn trailing_zeros(mut num: u32) -> u32 {
        if num == 0 {
            return 32;
        }

        let mut zeros = 0u32;
        while num & 1 == 0 {
            zeros += 1;
            num >>= 1;
        }
        zeros
    }

    fn to_cidr(ip: u32, bits: u32) -> String {
        format!("{}/{}", Self::int_to_ip(ip), bits)
    }

    pub fn ip_to_cidr(ip: String, mut n: i32) -> Vec<String> {
        let mut cidrs = Vec::new();
        let mut num = Self::ip_to_int(&ip);

        while n > 0 {
            let mut zeros = Self::trailing_zeros(num);

            zeros = zeros.min(31 - (n as u32).leading_zeros());

            let cnt = 1u32 << zeros;
            cidrs.push(Self::to_cidr(num, 32 - zeros));
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
