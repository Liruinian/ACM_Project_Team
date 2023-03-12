/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }
    var back int
    for x > back {
        back = x % 10 + back * 10
        x = x / 10 
    }
    if back == x || back/10 == x {
        return true
    }
    return false
}

// @lc code=end

