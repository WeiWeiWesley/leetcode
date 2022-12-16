package univaluedbinarytree

// CreateTree by array and sort by level-order
func CreateTree(input []int) *TreeNode {

	if len(input) == 0 {
		return &TreeNode{}
	}

	root := &TreeNode{Val: input[0]}

	//create a channel by input which without root
	nodes := make(chan int, len(input))

	for i := range input {
		//skip root
		if i == 0 {
			continue
		}

		nodes <- input[i]
	}
	close(nodes)

	root.create(nodes, -1)

	return root
}

func (root *TreeNode) create(nodes chan int, skip int) {

	if root == nil {
		return
	}

	currentNode := root
	queueNode := []*TreeNode{}

	for val := range nodes {

		if val != skip {
			newNode := &TreeNode{Val: val, Parent: currentNode}
			currentNode.Left = newNode
			queueNode = append(queueNode, newNode)
		}

		//取第二次
		val, ok := <-nodes
		if !ok {
			break
		}

		if val != skip {
			newNode := &TreeNode{Val: val, Parent: currentNode}
			currentNode.Right = newNode
			queueNode = append(queueNode, newNode)
		}

		currentNode = queueNode[0]
		queueNode = queueNode[1:]
	}
}

func (n *TreeNode) insert(val int) {

	if n == nil {
		return
	}

	currentNode := n
	queueNode := []*TreeNode{}

	for {
		if currentNode.Left == nil {
			currentNode.Left = &TreeNode{Val: val}
			return
		} else {
			queueNode = append(queueNode, currentNode.Left)
		}

		if currentNode.Right == nil {
			currentNode.Right = &TreeNode{Val: val}
			return
		} else {
			queueNode = append(queueNode, currentNode.Right)
		}

		currentNode = queueNode[0]
		queueNode = queueNode[1:]
	}
}
