/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	sx := string(x)
	sxarr := strings.Split(sx, "")
	outp := ""
	negative := 0
	if sxarr[0] == "-"{
		negative = 1
		sxarr =sxarr[1:]
	}
	for i := len(sxarr) - 1; i >= 0; i--{
		outp = outp + sxarr[i]
	}
	outpI,_ := strconv.ParseInt(outp,10,0)
	if negative == 1 {
		outpI = 0 - outpI
	}
	return int(outpI)
}
// @lc code=end

