/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	l := 0
	r := 9
	m := map[string]int
	arrs := strings.Split(s,"")
	for i := 0; i < len(arrs) - 8; i++ {
		m[arrs[i]] += 1 
	}
	var ans []string
	for k, v := range m {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}
// @lc code=end

