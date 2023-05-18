package main

import (
	"fmt"
)

// 節點結構
type Node struct {
	value int
	next  *Node
}

// 佇列結構
type Queue struct {
	front *Node // 前端節點
	back  *Node // 後端節點
	size  int   // 佇列大小
}

// 入隊操作
func (q *Queue) push(val int) {
	newNode := &Node{
		value: val,
		next:  nil,
	}

	if q.front == nil {
		q.front = newNode
		q.back = newNode
	} else {
		q.back.next = newNode
		q.back = newNode
	}

	q.size++
}

// 出隊操作
func (q *Queue) pop() {
	if q.front == nil {
		return // 佇列為空，無法出隊
	}

	//FIFO
	q.front = q.front.next
	q.size--

	if q.front == nil {
		q.back = nil // 若前端為空，則後端也為空
	}
}

// 獲取前端元素
func (q *Queue) getfront() int {
	if q.front == nil {
		return -1 // 佇列為空，返回特定值
	}

	return q.front.value
}

// 獲取後端元素
func (q *Queue) getback() int {
	if q.back == nil {
		return -1 // 佇列為空，返回特定值
	}

	return q.back.value
}

// 判斷佇列是否為空
func (q *Queue) isempty() bool {
	return q.front == nil
}

// 獲取佇列大小
func (q *Queue) qsize() int {
	return q.size
}

func main() {
	queue := Queue{} // 創建一個新的佇列

	// 測試各種操作
	queue.push(1)
	queue.push(2)
	queue.push(3)

	fmt.Println("Front element:", queue.getfront()) // 輸出: 1
	fmt.Println("Back element:", queue.getback())   // 輸出: 3
	fmt.Println("Is empty:", queue.isempty())       // 輸出: false
	fmt.Println("Queue size:", queue.qsize())       // 輸出: 3

	queue.pop()
	fmt.Println("Front element after pop:", queue.getfront()) // 輸出: 2

	queue.pop()
	queue.pop()
	fmt.Println("Is empty after pops:", queue.isempty()) // 輸出: true
	fmt.Println("Queue size after pops:", queue.qsize()) // 輸出: 0
}
