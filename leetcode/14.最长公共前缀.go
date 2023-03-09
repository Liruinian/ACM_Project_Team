/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 || strs[0] == "" {
		return ""
	}
	i := 0
	var char, ans string
	for {
		for j, st := range strs {
			starr := strings.Split(st, "")
			if i == len(starr) {
				return ans
			}
			if j == 0 {
				char = starr[i]
			} else if starr[i] != char {
				return ans
			}

		}
		i++
		ans += char
	}
}

// @lc code=end

