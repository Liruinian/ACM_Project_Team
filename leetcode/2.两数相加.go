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
	ans := new(ListNode)
	pre := ans
	last := ans
	carry := 0
	for l1 != nil || l2 != nil{
		if l1 == nil {
			l1 = new(ListNode)
		}else if l2 == nil{
			l2 = new(ListNode)
		}
			n := l1.Val + l2.Val
			if carry == 1{
				n++;
				carry--
			}
			tempn := new(ListNode)
			if n > 9{
				carry++
				tempn.Val = n - 10
			}else {
				tempn.Val = n
			}
			
			pre.Next = tempn
			pre = tempn
			last = tempn
			l1 = l1.Next
			l2 = l2.Next
		
		
	}
	if carry == 1{
		addOne := new(ListNode)
		addOne.Val = 1
		last.Next = addOne
	}
	return ans.Next
}

// @lc code=end

