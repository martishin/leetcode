package ip_to_cidr_751

import (
	"fmt"
	"strconv"
	"strings"
)

func toBinary(ip int) string {
	var sb strings.Builder
	for i := 31; i >= 0; i-- {
		bit := (ip >> i) & 1
		sb.WriteByte('0' + byte(bit))
		if i%8 == 0 && i != 32 {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}

func ipToInt(ip string) int {
	result := 0
	for _, part := range strings.Split(ip, ".") {
		octet, _ := strconv.Atoi(part)

		result = (result << 8) + octet
	}

	return result
}

func toCIDR(ip, bits int) string {
	return toIp(ip) + "/" + fmt.Sprint(bits)
}

func toIp(ip int) string {
	var sb strings.Builder

	for i := 3; i >= 0; i-- {
		octet := strconv.Itoa(ip >> (i * 8) & 0xFF) // 0b11111111 or 255
		sb.WriteString(octet)
		if i > 0 {
			sb.WriteByte('.')
		}
	}
	return sb.String()
}

func trailingZeros(num int) int {
	if num == 0 {
		return 32
	}

	zeros := 0
	for num&1 == 0 {
		zeros++
		num >>= 1
	}
	return zeros
}

func ipToCIDR(ip string, n int) []string {
	var cidrs []string

	num := ipToInt(ip)
	for n > 0 {
		zeros := trailingZeros(num)

		for 1<<zeros > n {
			zeros--
		}

		cnt := 1 << zeros
		cidrs = append(cidrs, toCIDR(num>>zeros<<zeros, 32-zeros))
		num += cnt
		n -= cnt
	}

	return cidrs
}

func Test() {
	fmt.Println(ipToCIDR("255.0.0.7", 10))
}
