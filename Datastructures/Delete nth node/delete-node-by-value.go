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
	delete(11)
	fmt.Println()
	print()
}

func delete(val int) {
	cur := head.next
	pre := head
	if pre.val == val {
		head = head.next
		return
	}

	for cur != nil {
		if cur.val == val {
			pre.next = pre.next.next
		}
		cur = cur.next
		pre = pre.next
	}
}

func insert(x int) {
	temp := &node{
		val:  x,
		next: nil,
	}
	if head == nil {
		head = temp
		return
	}
	cur := head
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
