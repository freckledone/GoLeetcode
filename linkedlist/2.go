package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode = &ListNode{0, nil}

	var head *ListNode = l3
	*head = *l3

	adder := 0

	for l1 != nil || l2 != nil || adder > 0 {
		value1, value2 := 0, 0

		if l1 != nil {
			value1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			value2 = l2.Val
			l2 = l2.Next
		}

		val := value1 + value2 + adder

		adder = val / 10

		val = val % 10

		l3.Next = &ListNode{val, nil}
		l3 = l3.Next
	}

	return head.Next
}
