/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	var lclose []string
	sarr := strings.Split(s, "")
	for i := 0; i < len(sarr); i++ {
		switch sarr[i] {
		case "(":
			lclose = append(lclose, "a")
		case "[":
			lclose = append(lclose, "b")
		case "{":
			lclose = append(lclose, "c")
		case ")":
			if len(lclose) == 0 {
				return false
			}
			if lclose[len(lclose)-1] != "a" {
				return false
			} else {
				lclose = lclose[:len(lclose)-1]
			}
		case "]":
			if len(lclose) == 0 {
				return false
			}
			if lclose[len(lclose)-1] != "b" {
				return false
			} else {
				lclose = lclose[:len(lclose)-1]
			}
		case "}":
			if len(lclose) == 0 {
				return false
			}
			if lclose[len(lclose)-1] != "c" {
				return false
			} else {
				lclose = lclose[:len(lclose)-1]
			}
		default:
			break
		}
	}
	if len(lclose) == 0 {
		return true
	} else {
		return false
	}
}

// @lc code=end

