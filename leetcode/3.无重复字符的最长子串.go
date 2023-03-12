/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var ansarr []string
	sarr := strings.Split(s, "")
	for _,schar := range sarr {
		for j := 0; j < len(ansarr);j++{
			if ansarr[j] == schar {
                ansarr = ansarr[:j]
                break
            }else{
				ansarr = append(ansarr, schar)
			}
		}
	}
	return len(ansarr)
}
// @lc code=end

