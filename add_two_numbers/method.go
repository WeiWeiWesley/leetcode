package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	resultNode := &ListNode{}
	tmp := resultNode
	previous := 0 //進位

	for {

		//不等長 list
		if l1 == nil && l2 == nil {
			break
		}

		sum := 0

		sum += previous //前次進位

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		currentVal := sum % 10

		//新進位
		previous = int(sum / 10)

		tmp.Next = &ListNode{
			Val: currentVal,
		}

		tmp = tmp.Next
	}

	//檢查最大位數是否進位
	if previous > 0 {
		tmp.Next = &ListNode{
			Val: previous,
		}
	}

	return resultNode.Next
}
