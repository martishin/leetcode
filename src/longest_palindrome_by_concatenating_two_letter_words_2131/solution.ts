function longestPalindrome(words: string[]): number {
  const wordToCount = new Map<string, number>();

  for (const word of words) {
    wordToCount.set(word, (wordToCount.get(word) || 0) + 1);
  }

  const visited = new Set<string>();
  let usedCenter = false;
  let result = 0;

  for (const [word, count] of wordToCount.entries()) {
    if (visited.has(word)) {
      continue;
    }

    visited.add(word);

    if (word[0] === word[1]) {
      if (count % 2 === 0) {
        result += count * 2;
      } else {
        result += (count - 1) * 2;
        if (!usedCenter) {
          usedCenter = true;
          result += 2;
        }
      }
    } else {
      const pair = word[1] + word[0];
      const pairCount = wordToCount.get(pair) || 0;

      if (pairCount > 0) {
        visited.add(pair);
        result += 2 * 2 * Math.min(count, pairCount);
      }
    }
  }

  return result;
}

export function test() {
  console.log(longestPalindrome(["lc", "cl", "gg"]));
  console.log(longestPalindrome(["ab", "ty", "yt", "lc", "cl", "ab"]));
  console.log(longestPalindrome(["cc", "ll", "xx"]));
}
