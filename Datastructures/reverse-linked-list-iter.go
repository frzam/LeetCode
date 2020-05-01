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
	reverse()
	print()
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
	cur := head
	for cur != nil {
		fmt.Printf("%d ", cur.val)
		cur = cur.next
	}
	fmt.Println()
}

func reverse() {
	var prev *node
	cur := head
	var next *node
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next

	}
	head = prev
}
