package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1234
// 2134
// 3214
// 4321
func reverseList(head *ListNode) *ListNode {

	var previous *ListNode

	current := head

	for current != nil {

		tmp := current.Next

		current.Next = previous
		previous = current
		current = tmp

	}

	return previous
}
