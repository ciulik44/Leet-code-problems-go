package main

import "fmt"

func twoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}

func main() {
	target := 9
	nums := [5]int{1, 2, 3, 4, 5}
	result := twoSum(nums[:], target)
	fmt.Println(result)
}
