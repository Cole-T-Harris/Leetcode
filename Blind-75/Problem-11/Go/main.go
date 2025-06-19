package main

import "fmt"

func maxArea(height []int) int {
    maxContainer := 0
	l := 0
	r := len(height) - 1
	for l < r {
		maxContainer = max(maxContainer, min(height[l], height[r]) * (r - l))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxContainer
}

func main() {
	heights := []int{1,8,6,2,5,4,8,3,7}
	fmt.Println(maxArea(heights))
}