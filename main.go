package main

import "fmt"

func main() {
	l1 := ConvertToLinkedList(243)
	l2 := ConvertToLinkedList(564)

	result := AddTwoNumbers(l1, l2)
	result = Reverse(result)

	for result != nil {
		fmt.Print(result.Val)
		result = result.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
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

func ConvertToLinkedList(num int) *ListNode {
	current := &ListNode{}
	head := current

	for num > 0 {
		current.Next = &ListNode{Val: num % 10}
		current = current.Next
		num /= 10
	}

	return head.Next
}

// Сложность О(n)
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int = 0

	var currentNode *ListNode = &ListNode{}
	var head *ListNode = currentNode

	for l1 != nil && l2 != nil {

		sum := l1.Val + l2.Val + carry

		carry = sum / 10

		currentNode.Next = &ListNode{Val: sum % 10}
		currentNode = currentNode.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	carry, currentNode = Fill(l1, currentNode, carry)
	carry, currentNode = Fill(l2, currentNode, carry)

	if carry > 0 {
		currentNode.Next = &ListNode{Val: carry}
	}

	return head.Next
}

func Fill(list *ListNode, cur *ListNode, carry int) (int, *ListNode) {
	for list != nil {
		sum := list.Val + carry
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
		list = list.Next
	}
	return carry, cur
}
