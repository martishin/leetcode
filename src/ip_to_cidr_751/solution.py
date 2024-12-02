from typing import List


class Solution:
    def ip_to_int(self, ip: str):
        result = 0

        for part in ip.split("."):
            octet = int(part)
            result = (result << 8) + octet

        return result

    def trailing_zeros(self, num: int):
        if num == 0:
            return 32

        zeros = 0
        while num & 1 == 0:
            zeros += 1
            num >>= 1

        return zeros

    def to_cidr(self, ip: int, bits: int) -> str:
        print(bits)
        return f"{self.to_ip(ip)}/{bits}"

    def to_ip(self, ip: int) -> str:
        octets = []
        for i in range(3, -1, -1):
            octet = (ip >> (i * 8)) & 0xFF
            octets.append(str(octet))
        return ".".join(octets)

    def ipToCIDR(self, ip: str, n: int) -> List[str]:
        cidrs = []
        num = self.ip_to_int(ip)
        while n > 0:
            zeros = self.trailing_zeros(num)

            print(zeros)
            while (1 << zeros) > n:
                zeros -= 1

            cnt = 1 << zeros
            print(zeros)
            cidrs.append(self.to_cidr(num, 32 - zeros))
            num += cnt
            n -= cnt

        return cidrs


def test():
    solution = Solution()
    print(solution.ipToCIDR("255.0.0.7", 10))
    print(solution.ipToCIDR("117.145.102.62", 8))
    print(solution.ipToCIDR("0.0.0.0", 1))
