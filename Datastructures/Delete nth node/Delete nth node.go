package main

import "fmt"

type node struct {
	val  int
	next *node
}

var head *node

func main() {
	insert(2)
	insert(4)
	insert(6)
	insert(8)
	print()
	delete(3)
	fmt.Println()
	print()
}
func delete(n int) {
	if n == 1 {
		head = head.next
		return
	}
	temp := head
	for temp != nil && n > 2 {
		temp = temp.next
		n--
	}
	temp.next = temp.next.next
}

func insert(x int) {
	cur := head
	temp := &node{
		val:  x,
		next: nil,
	}
	if head == nil {
		head = temp
		return
	}
	for cur.next != nil {
		cur = cur.next
	}

	cur.next = temp
}

func print() {
	temp := head
	for temp != nil {
		fmt.Printf("%d ", temp.val)
		temp = temp.next
	}
}
