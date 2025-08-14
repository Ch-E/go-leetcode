package main

import "fmt"

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}
	printList(head)

	fmt.Println("\nDuplicates deleted:")
	dedupHead := deleteDuplicates(head)
	printList(dedupHead)
}

// Given the head of a sorted linked list, delete all duplicates such that each element appears only once. Return the linked list sorted as well.

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}
