package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

var head *Node

func main() {
	fmt.Printf("Insert the number of nodes to insert.\n")
	var n int
	var x int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Printf("Enter the number \n")
		fmt.Scanf("%d\n", &x)
		Insert(x)
		Print()
	}
}

func Insert(x int) {
	temp := &Node{
		Val:  x,
		Next: head,
	}
	head = temp
}

func Print() {
	temp := head
	fmt.Println("Printing the list:")
	for temp != nil {
		fmt.Printf(" %d", temp.Val)
		temp = temp.Next
	}
	fmt.Println()
}
