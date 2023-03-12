/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	var ans int
	arrs := strings.Split(s, "")
	l := len(arrs) - 1
	for i := 0; i < len(arrs); i++ {
		a := arrs[i]
		if a == "I" && i < l {
			if string(arrs[i+1]) == "V" {
				ans += 4
				i += 1
				continue
			}
			if string(arrs[i+1]) == "X" {
				ans += 9
				i += 1
				continue
			}

		} else if a == "X" && i < l {
			if string(arrs[i+1]) == "L" {
				ans += 40
				i += 1
				continue
			}
			if string(arrs[i+1]) == "C" {
				ans += 90
				i += 1
				continue
			}

		} else if a == "C" && i < l {
			if string(arrs[i+1]) == "D" {
				ans += 400
				i += 1
				continue
			}
			if string(arrs[i+1]) == "M" {
				ans += 900
				i += 1
				continue
			}

		}
		switch a {
		case "I":
			ans += 1
		case "V":
			ans += 5
		case "X":
			ans += 10
		case "L":
			ans += 50
		case "C":
			ans += 100
		case "D":
			ans += 500
		case "M":
			ans += 1000
		default:
			ans += 0
		}

	}
	return ans
}

// @lc code=end

