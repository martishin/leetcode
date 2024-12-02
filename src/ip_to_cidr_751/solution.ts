function ipToCIDR(ip: string, n: number): string[] {
  const ipToInt = (ip: string): number => {
    let result = 0;
    for (const part of ip.split(".")) {
      const octet = parseInt(part, 10);
      result = (result << 8) + octet;
    }
    return result;
  };

  const trailingZeros = (num: number): number => {
    if (num === 0) {
      return 32;
    }

    let zeros = 0;
    while ((num & 1) === 0) {
      zeros++;
      num >>= 1;
    }
    return zeros;
  };

  const toCIDR = (ip: number, bits: number): string => {
    return `${toIp(ip)}/${bits}`;
  };

  const toIp = (ip: number): string => {
    const octets: string[] = [];
    for (let i = 3; i >= 0; i--) {
      const octet = (ip >> (i * 8)) & 0xff;
      octets.push(octet.toString());
    }
    return octets.join(".");
  };

  const powerOfTwo = (exponent: number): number => {
    return Math.pow(2, exponent);
  };

  const cidrs: string[] = [];
  let num = ipToInt(ip);

  while (n > 0) {
    let zeros = trailingZeros(num);

    while (powerOfTwo(zeros) > n) {
      zeros--;
    }

    const cnt = 1 << zeros;
    cidrs.push(toCIDR(num, 32 - zeros));
    num += cnt;
    n -= cnt;
  }

  return cidrs;
}

export function test(): void {
  console.log(ipToCIDR("255.0.0.7", 10));
  console.log(ipToCIDR("117.145.102.62", 8));
  console.log(ipToCIDR("0.0.0.0", 1));
}
