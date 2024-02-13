package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// Original ListNode
	prehead := &ListNode{Val: -1, Next: nil}

	// Copying the value
	copiedValue := *prehead
	copiedValue.Val = 10

	// Output
	fmt.Println("Original ListNode:", prehead.Val)
	fmt.Println("Copied value:", copiedValue.Val)
		
	// Copying the pointer
	copiedPointer := prehead
	copiedPointer.Val = 20

	// Output
	fmt.Println("Original ListNode:", prehead.Val)
	fmt.Println("Copied pointer:", copiedPointer.Val)
}