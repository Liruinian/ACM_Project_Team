package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
}
func search(nums []int, target int) int {
	if nums[len(nums)/2] > target {
		for i := len(nums)/2 - 1; i >= 0; i-- {
			if nums[i] == target {
				return i
			}
		}
	} else {
		for i := len(nums) / 2; i <= len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
	}
	return -1
}

func isBadVersion(version int) bool {
	if version == 1 {
		return true
	} else {
		return false
	}
}
