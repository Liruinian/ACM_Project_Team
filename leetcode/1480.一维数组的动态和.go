/*
 * @lc app=leetcode.cn id=1480 lang=golang
 *
 * [1480] 一维数组的动态和
 */

// @lc code=start
func runningSum(nums []int) []int {
    ans := nums
    for a, b := range nums{
        if a != 0{
            ans[a] = ans[a-1] + b
        }
    }
    return ans
}
// @lc code=end

