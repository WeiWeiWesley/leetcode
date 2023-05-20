package main

import "fmt"

type Hashmap struct {
	bucket [4]*Node
	len    int
}

type Node struct {
	key  string
	val  int
	next *Node
}

func (h *Hashmap) hash(k string) int {
	// 先把string轉成 ASCII code再 % 桶子的大小
	return int(k[0]) % len(h.bucket)
}

func (h *Hashmap) set(key string, val int) {

	//create hash key
	hashKey := h.hash(key)

	//不存在直接新增
	if h.bucket[hashKey] == nil {
		h.bucket[hashKey] = &Node{
			key: key,
			val: val,
		}

		h.len++
		return
	}

	node := h.bucket[hashKey]

	for node != nil {
		//替換相同 key 裡面的值
		if node.key == key {
			node.val = val
			return
		}

		node = node.next
	}

	//沒找到一樣的key, 新節點放在最前面
	h.bucket[hashKey] = &Node{
		key:  key,
		val:  val,
		next: h.bucket[hashKey],
	}

	h.len++

	return
}

func (h *Hashmap) get(key string) (int, bool) {

	hashKey := h.hash(key)

	//不存在
	if h.bucket[hashKey] == nil {
		return 0, false
	}

	node := h.bucket[hashKey]

	for node != nil {
		//存在回傳值
		if node.key == key {
			return node.val, true
		}

		node = node.next
	}

	//不存在
	return 0, false
}

func (h *Hashmap) delete(key string) {

	hashKey := h.hash(key)

	//不存在
	if h.len == 0 {
		return
	}

	node := h.bucket[hashKey]
	previous := node //刪除需要用前接後，故需記錄前一節點

	//若為第一個節點
	if node.key == key {
		h.bucket[hashKey] = nil
		h.len--
		return
	}

	for node != nil {
		if node.key == key {
			previous.next = node.next //前接後
			h.len--
			return
		}

		previous = node  //當前節點圍下一輪節的 pre
		node = node.next //移動
	}

}

func main() {

	h := Hashmap{}

	h.set("dog", 10)

	fmt.Println(h.get("dog"))

	h.set("dog", 20)
	h.set("apple", 11)
	h.set("dif", 12)

	fmt.Println(h.get("dog"))
	fmt.Println(h.get("apple"))
	fmt.Println(h.get("dif"))

	h.delete("dif")
	fmt.Println(h.get("dif"))
}
