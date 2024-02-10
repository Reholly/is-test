package main

import "fmt"

func main() {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 1}
	l2.Next = &ListNode{Val: 9}

	list := addTwoNumbers(l1, l2)
	list = Reverse(list)

	for list != nil {
		fmt.Print(list.Val)
		list = list.Next
	}
}

func Reverse(list *ListNode) *ListNode {
	var previousNode *ListNode

	currentNode := list
	nextNode := currentNode.Next

	for nextNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}

	return previousNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// Сложность О(n)
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0

	var cur *ListNode = &ListNode{}
	var head *ListNode = cur

	for l1 != nil && l2 != nil {

		sum := l1.Val + l2.Val + carry

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	carry, cur = fill(l1, carry, cur)
	carry, cur = fill(l2, carry, cur)

	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}

	return head.Next
}

func fill(l1 *ListNode, carry int, cur *ListNode) (int, *ListNode) {
	for l1 != nil {
		sum := l1.Val + carry
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
		l1 = l1.Next
	}
	return carry, cur
}
