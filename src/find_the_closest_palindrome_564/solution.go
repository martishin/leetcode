package find_the_closest_palindrome_564

import (
	"math"
	"strconv"
)

func nearestPalindromic(numberStr string) string {
	number, _ := strconv.ParseInt(numberStr, 10, 64)
	if number <= 10 {
		return strconv.FormatInt(number-1, 10)
	}
	if number == 11 {
		return "9"
	}

	length := len(numberStr)
	leftHalf, _ := strconv.ParseInt(numberStr[:(length+1)/2], 10, 64)

	palindromeCandidates := make([]int64, 5)
	palindromeCandidates[0] = generatePalindromeFromLeft(leftHalf-1, length%2 == 0)
	palindromeCandidates[1] = generatePalindromeFromLeft(leftHalf, length%2 == 0)
	palindromeCandidates[2] = generatePalindromeFromLeft(leftHalf+1, length%2 == 0)
	palindromeCandidates[3] = int64(math.Pow10(length-1)) - 1
	palindromeCandidates[4] = int64(math.Pow10(length)) + 1

	var nearestPalindrome int64
	minDifference := int64(math.MaxInt64)

	for _, candidate := range palindromeCandidates {
		if candidate == number {
			continue
		}
		difference := abs(candidate - number)
		if difference < minDifference || (difference == minDifference && candidate < nearestPalindrome) {
			minDifference = difference
			nearestPalindrome = candidate
		}
	}

	return strconv.FormatInt(nearestPalindrome, 10)
}

func generatePalindromeFromLeft(leftHalf int64, isEvenLength bool) int64 {
	palindrome := leftHalf
	if !isEvenLength {
		leftHalf /= 10
	}
	for leftHalf > 0 {
		palindrome = palindrome*10 + leftHalf%10
		leftHalf /= 10
	}
	return palindrome
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
