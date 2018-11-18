package main

import "fmt"

// 粗暴遍历: O(n^2)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target = 1

	result := twoSum(nums, target)
	fmt.Printf("Result: %+v", result)
}
