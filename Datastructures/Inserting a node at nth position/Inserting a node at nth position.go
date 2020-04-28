package main

import "fmt"

type node struct {
	val  int
	next *node
}

func main() {
	insert(2, 1) // 2
	insert(3, 2) // 2 3
	insert(4, 1) // 4 2 3
	insert(5, 2) // 4 5 2 3
	print()
}

var head *node

func insert(x, pos int) {

	top := head
	if pos == 1 {
		temp := &node{
			val:  x,
			next: head,
		}
		head = temp
		return
	}
	for i := 1; i < pos-1; i++ {
		top = top.next
	}
	temp := &node{
		val:  x,
		next: top.next,
	}
	top.next = temp
}
func print() {
	temp := head
	fmt.Println("Printing the list.")
	for temp != nil {
		fmt.Printf(" %d", temp.val)
		temp = temp.next
	}
	fmt.Println()
}
