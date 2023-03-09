/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	cal := new(ListNode)
	ans := cal
	carry := 0
	for l1 != nil || l2 != nil {
		cal.Val = l1.Val + l2.Val
		if carry == 1 {
			cal.Val++
			carry = 0
		}
		if cal.Val > 9 {
			carry = 1
			cal.Val -= 10
		}
		if l1 == nil{
			l2 = l2.Next
		}else if l2 == nil{
			l1 = l1.Next
		}else{
			l1 = l1.Next
			l2 = l2.Next
		}
		if l1 == nil && l2 == nil {
			break
		} else {
			next := new(ListNode)
			cal.Next = next
			cal = next
		}

	}
	cal.Next = nil
	return ans
}

// @lc code=end

