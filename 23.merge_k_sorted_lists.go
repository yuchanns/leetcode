/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	left, right := make([]*ListNode, 0), make([]*ListNode, 0)
	var cmp *ListNode
	for i := range lists {
		if lists[i] != nil {
			cmp = lists[i]
			break
		}
	}
	if cmp == nil {
		return nil
	}
	for cmp.Next != nil {
		left = append(left, cmp)
		cmp = cmp.Next
	}
	for i := 0; i < k; i++ {
		toCmp := lists[i]
		if toCmp == nil {
			continue
		}
		for {
			if toCmp.Val <= cmp.Val {
				left = append(left, toCmp)
			} else {
				right = append(right, toCmp)
			}
			if toCmp.Next == nil {
				break
			}
			toCmp = toCmp.Next
		}
	}
	sorted := append(qs(left), qs(right)...)
	returnNode := sorted[0]
	currentNode := returnNode
	for i := 1; i < len(sorted); i++ {
		currentNode.Next = sorted[i]
		currentNode = sorted[i]
		if i == len(sorted)-1 {
			currentNode.Next = nil
		}
	}
	return returnNode
}

func qs(nodes []*ListNode) []*ListNode {
	if len(nodes) <= 1 {
		return nodes
	}
	cmp := nodes[0]
	left, right := make([]*ListNode, 0), make([]*ListNode, 0)
	for i := 1; i < len(nodes); i++ {
		if nodes[i].Val <= cmp.Val {
			left = append(left, nodes[i])
		} else {
			right = append(right, nodes[i])
		}
	}
	return append(append(qs(left), cmp), qs(right)...)
}
