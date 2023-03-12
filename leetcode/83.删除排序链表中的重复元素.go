/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tempVal := head.Val
	node := head
	for node.Next != nil && node.Next.Next != nil {
		if node.Val == tempVal {
			node.Next = node.Next.Next
		}

		node = node.Next
		tempVal = node.Val
	}
	return head
}

// @lc code=end

