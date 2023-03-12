/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	temp := 1
	l := len(nums)
	for i := 0; i < l-1; i++ {
		if nums[i] != nums[i+1] {
			nums[temp] = nums[i+1]
			temp++
		}
	}
	return temp
}

// @lc code=end

