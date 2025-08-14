package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}

	printList(l1)
	printList(l2)
	printList(addTwoNumbers(l1, l2))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order, and each of their nodes contains a single digit.
Add the two numbers and return the sum as a linked list.
*/

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	sum := head

	var l1str strings.Builder
	var l2str strings.Builder

	for l1 != nil {
		l1str.WriteString(strconv.Itoa(l1.Val))
		l1 = l1.Next
	}

	for l2 != nil {
		l2str.WriteString(strconv.Itoa(l2.Val))
		l2 = l2.Next
	}

	l1fstr := l1str.String()
	l2fstr := l2str.String()

	l1runes := []rune(l1fstr)
	l2runes := []rune(l2fstr)

	slices.Reverse(l1runes)
	slices.Reverse(l2runes)

	l1reversed := string(l1runes)
	l2reversed := string(l2runes)

	// encountered error: strconv.Atoi: parsing "1000000000000000000000000000001": value out of range <nil>
	l1int, err1 := strconv.Atoi(l1reversed)
	l2int, err2 := strconv.Atoi(l2reversed)

	addNrs := l1int + l2int
	addNrsStr := strconv.Itoa(addNrs)
	addNrsRune := []rune(addNrsStr)
	slices.Reverse(addNrsRune)

	for i := range addNrsRune {
		sum.Val = int(addNrsRune[i] - '0')
		if i != len(addNrsRune)-1 {
			sum.Next = &ListNode{}
			sum = sum.Next
		}
	}

	if err1 != nil || err2 != nil {
		fmt.Printf("encountered error: %v %v\n", err1, err2)
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
