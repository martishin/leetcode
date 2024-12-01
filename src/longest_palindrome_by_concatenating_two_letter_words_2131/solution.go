package longest_palindrome_by_concatenating_two_letter_words_2131

import "fmt"

func longestPalindrome(words []string) int {
	wordToCount := make(map[string]int)

	for _, word := range words {
		wordToCount[word]++
	}

	putOneInCenter := false
	result := 0
	visited := make(map[string]struct{})

	for word, count := range wordToCount {
		if _, found := visited[word]; found {
			continue
		}

		visited[word] = struct{}{}

		if word[0] == word[1] {
			if count%2 == 0 {
				result += count * 2
			} else {
				result += (count - 1) * 2
				if !putOneInCenter {
					putOneInCenter = true
					result += 2
				}
			}
		} else {
			pair := string(word[1]) + string(word[0])
			if pairCount, found := wordToCount[pair]; found {
				visited[pair] = struct{}{}

				result += 2 * 2 * min(count, pairCount)
			}
		}
	}

	return result
}

func Test() {
	fmt.Println(longestPalindrome([]string{"lc", "cl", "gg"}))
	fmt.Println(longestPalindrome([]string{"ab", "ty", "yt", "lc", "cl", "ab"}))
	fmt.Println(longestPalindrome([]string{"cc", "ll", "xx"}))
}
