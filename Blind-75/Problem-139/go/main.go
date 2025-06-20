package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	if len(s) < 1 {
		return false
	}
    wordSet := make(map[string] struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}
	condition := make([]bool, len(s) + 1)
	condition[0] = true

	for i := 1; i < len(condition); i++ {
		for j := 0; j < i; j++ {
			if _, found := wordSet[s[j:i]]; found && condition[j] {
				condition[i] = true
			}
		}
	}
	return condition[len(s)]
}

func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))
}