package main

func twoSum(nums []int, target int) []int {
    nums_map := make(map[int]int)
	for i, num := range nums {
		nums_map[num] = i
	}
	for i, num := range nums {
		complement := target - num
		if index, found := nums_map[complement]; found && index != i {
			return []int{i, index}
		}
	}
	return []int{}
}

func main() {
	// Example usage
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	// Output: [0, 1]
	println(result[0], result[1])
}