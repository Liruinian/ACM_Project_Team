/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	digits[len(digits)-1] += 1
	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i-1]++
			digits[i] = 0
		}
	}
	if digits[0] > 9 {
		one := []int{1}
		digits[0] = 0
		digits = append(one[:], digits...)
	}
	return digits
}

// @lc code=end

