/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	arrs := strings.Split(s, " ")
	for i := len(arrs) - 1; i >= 0; i-- {
		if arrs[i] == "" {
			continue
		} else {
			wordarr := strings.Split(arrs[i], "")
			return len(wordarr)
		}
	}
	return 0
}

// @lc code=end

