/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		result  *ListNode
		current *ListNode
	)
	carry := 0
	for {
		if l1 == nil && l2 == nil && carry == 0 {
			break
		}
		if current == nil {
			current = &ListNode{
				Val: carry,
			}
		} else {
			current.Next = &ListNode{
				Val: carry,
			}
			current = current.Next
		}
		if l1 != nil {
			current.Val = (current.Val + l1.Val) % 10
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			current.Val = (current.Val + l2.Val) % 10
			carry += l2.Val
			l2 = l2.Next
		}
		carry /= 10
		if result == nil {
			result = current
		}
	}
	return result
}
