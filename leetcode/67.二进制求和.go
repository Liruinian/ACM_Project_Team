/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 */

// @lc code=start
func addBinary(a string, b string) string {
	func bstringtoInt(s string) int{
		arrs := strings.Split(s, "")
		ret := 0
		for i, num := range arrs{
			if num == "1" {
				ret += int(math.Pow(2, float64(i)))
			}
		}
		return ret
	}

	func inttoBString(inp int) string{
		for 
		return ret
	}

	ab := bstringtoInt(a)
	bb := bstringtoInt(b)

	ans := ab + bb


	
	
}
// @lc code=end

