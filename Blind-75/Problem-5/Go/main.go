package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	palindromeMatrix := make([][]bool, n)
	for i := range(palindromeMatrix) {
		palindromeMatrix[i] = make([]bool, n)
	}
	for i := range(palindromeMatrix) {
		palindromeMatrix[i][i] = true
	}

	maxLength := 1
	start := 0
	for substringLength := 2; substringLength <= n; substringLength++ {
		for i := range n - substringLength + 1 {
			j := i + substringLength - 1
			if substringLength == 2 {
				palindromeMatrix[i][j] = s[i] == s[j]
			} else {
				palindromeMatrix[i][j] = s[i] == s[j] && palindromeMatrix[i+1][j-1]
			}
			if palindromeMatrix[i][j] && substringLength > maxLength {
				maxLength = substringLength
				start = i
			}
		}
	}
	return s[start:start + maxLength]	
}

func main() {
	s := "bb"
	fmt.Println(longestPalindrome(s))
}