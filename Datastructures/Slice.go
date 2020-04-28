package main

import (
	"fmt"
)

func main() {
	s := make(slice, 0)
	fmt.Println(s)
	c := count(s)
	fmt.Println(c)
	fmt.Println("isEmpty():", isEmpty(s))
	s = insert(s, 10)
	fmt.Println(s)
	c = count(s)
	fmt.Println(c)

	s = insert(s, 2)
	s = insert(s, 4)
	s = insert(s, 6)
	s = insert(s, 7)
	s = insert(s, 9)
	printSlice(s)
	s = insertPos(s, -1, 0)
	printSlice(s)
	s = insertPos(s, -5, 3)
	printSlice(s)
	s = remove(s, 3)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)
	s = remove(s, 0)
	printSlice(s)ku
	fmt.Println("c : ", count(s))
}

type slice []int

var end int

func init() {
	end = -1
}

func printSlice(s slice) {
	if len(s) == end {
		fmt.Print(s)
	} else {
		fmt.Println(s[:end+1])
	}
}

func count(s slice) int {
	return end + 1
}

func isEmpty(s slice) bool {
	return end == -1
}

func insert(s slice, e int) slice {
	if len(s) == count(s) {
		newSize := count(s)*3/2 + 1
		newLen := count(s)*3/2 + 1
		newSlice := make(slice, newLen, newSize)
		copy(newSlice, s)
		s = newSlice
	}
	s[count(s)] = e
	end++
	return s
}

func insertPos(s slice, e, pos int) slice {
	if len(s) == count(s) {
		newSize := count(s)*3/2 + 1
		newLen := count(s)*3/2 + 1
		newSlice := make(slice, newLen, newSize)
		copy(newSlice, s)
		s = newSlice
	}
	for i := len(s) - 2; i >= pos; i-- {
		s[i+1] = s[i]
	}
	s[pos] = e
	end++
	return s
}

func remove(s slice, pos int) slice {
	for i := pos; i < len(s)-1; i++ {
		s[i] = s[i+1]
	}
	end--
	return s
}
