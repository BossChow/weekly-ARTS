package main

import "fmt"

// 粗暴遍历: O(n^2)
func twoSumByIterating(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// 使用map: O(n)
func twoSumByMap(nums []int, target int) []int {
	mapping := make(map[int]int, 0)
	for i, val := range nums {
		mapping[val] = i
	}

	for i, val := range nums {
		complement := target - val
		if val, ok := mapping[complement]; ok {
			return []int{i, val}
		}
	}

	return nil
}

func twoSumByMap2(nums []int, target int) []int {
	mapping := make(map[int]int)
	for i, v := range nums {
		complement := target - v
		if val, ok := mapping[complement]; ok {
			return []int{i, val}
		}

		mapping[v] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	var target = 9

	result := twoSumByIterating(nums, target)
	fmt.Printf("Result: %+v\n", result)

	twoSumByMap(nums, target)
	fmt.Printf("Result: %+v\n", result)

	twoSumByMap2(nums, target)
	fmt.Printf("Result: %+v\n", result)
}
