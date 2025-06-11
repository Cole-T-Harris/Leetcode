package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	maxSubString := make(map[byte]struct{})
	maxLength := 0
	for right < len(s) {
		if _, found := maxSubString[s[right]]; found {
			delete(maxSubString, s[left])
			left++
			continue
		}
		maxSubString[s[right]] = struct{}{}
		right++
		if len(maxSubString) > maxLength {
			maxLength = len(maxSubString)
		}
	}
	return maxLength
}

func main() {
	s := "abcabcbb"
	d := "qrsvbspk"
	result := lengthOfLongestSubstring(s)
	resultd := lengthOfLongestSubstring(d)
	fmt.Println(result) // Should print 3 for "abc"
	fmt.Println(resultd)
}