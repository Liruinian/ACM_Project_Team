/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] N 字形变换
 */

// @lc code=start

import "strings"

func convert(s string, numRows int) string {
	var rowArr [][]string
	sArr := strings.Split(s, "")
	k := 0
	rev := 0
	for i := 0; i < len(sArr); i++ {
        if k == numRows - 1 {
			rev = 1
		}else if k == 0 && rev == 1{
			rev = 0
		}
		
		if rev == 0{
			k++
		}else{
			k--
		}
		
		rowArr[k] := append(rowArr,s[i])
		
		
    }
}
// @lc code=end

