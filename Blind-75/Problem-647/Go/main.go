package main

import (
	"fmt"
	"strings"
)

func countSubstrings(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var builder strings.Builder
	builder.WriteString("^")
	for _, ch := range s {
		builder.WriteString("#")
		builder.WriteRune(ch)
	}
	builder.WriteString("#$")
	transformedString := builder.String()
	pos, maxRight, palindromCount := 0, 0, 0
	palindromeRadii := make([]int, len(transformedString))
	for i := 1; i < len(palindromeRadii)-1; i++ {
		mirroredPosition := 2 * pos - i
		if i < maxRight {
			palindromeRadii[i] = min(maxRight - i, palindromeRadii[mirroredPosition])
		} else {
			palindromeRadii[i] = 1
		}
		for transformedString[i - palindromeRadii[i]] == transformedString[i + palindromeRadii[i]] {
			palindromeRadii[i]++
		}
		if i + palindromeRadii[i] > maxRight {
			maxRight = i + palindromeRadii[i]
			pos = i
		}
		palindromCount += int(palindromeRadii[i] / 2)
	}
	return palindromCount
}

func main() {
	fmt.Println(countSubstrings("abc"))
}