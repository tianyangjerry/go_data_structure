package linkedList

import (
	"cmp"
	"fmt"
)

// DoublyLinkedListNode 双向链表节点
type DoublyLinkedListNode[T cmp.Ordered] struct {
	Value T
	Next  *DoublyLinkedListNode[T]
	Prev  *DoublyLinkedListNode[T]
}

// DoublyLinkedList 双向链表
type DoublyLinkedList[T cmp.Ordered] struct {
	Head   *DoublyLinkedListNode[T]
	Tail   *DoublyLinkedListNode[T]
	Length int
}

func (list *DoublyLinkedList[T]) GetSize() int {
	return list.Length
}

func (list *DoublyLinkedList[T]) Traverse() {
	head := list.Head

	for head.Next != nil {
		fmt.Printf("%v -> ", head.Value)
		head = head.Next
	}
	println()
}

func (list *DoublyLinkedList[T]) ReverseTraverse() {
	tail := list.Tail

	for tail.Prev != nil {
		fmt.Printf("%v -> ", tail.Value)
		tail = tail.Prev
	}
	println()
}

func (list *DoublyLinkedList[T]) PushBack(value T) {
	newNode := &DoublyLinkedListNode[T]{
		Value: value,
	}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Tail.Next = newNode
		newNode.Prev = list.Tail
		list.Tail = newNode
	}

	list.Length++
}

func (list *DoublyLinkedList[T]) PopBack() T {
	if list.Length == 0 {
		panic("List is empty")
	}

	value := list.Tail.Value
	list.Tail = list.Tail.Prev
	list.Tail.Next = nil
	list.Length--
	return value
}

func (list *DoublyLinkedList[T]) PushFront(value T) {
	newNode := &DoublyLinkedListNode[T]{
		Value: value,
	}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		list.Head.Prev = newNode
		newNode.Next = list.Head
		list.Head = newNode
	}

	list.Length++
}

func (list *DoublyLinkedList[T]) PopFront() T {
	if list.Length == 0 {
		panic("List is empty")
	}

	value := list.Head.Value
	list.Head = list.Head.Next
	list.Head.Prev = nil
	list.Length--
	return value
}

func (list *DoublyLinkedList[T]) GetIndex(index int) (*DoublyLinkedListNode[T], error) {
	if index >= list.Length || index < 0 {
		return nil, fmt.Errorf("index out of range: %d", index)
	}

	node := list.Head
	for i := 0; i < index; i++ {
		node = node.Next
	}

	return node, nil
}

func (list *DoublyLinkedList[T]) InsertAfter(n1 *DoublyLinkedListNode[T], value T) {
	newNode := &DoublyLinkedListNode[T]{
		Value: value,
	}
	if n1.Next != nil {
		n1.Next.Prev = newNode
		newNode.Next = n1.Next
	} else {
		list.Tail = newNode
	}

	n1.Next = newNode
	newNode.Prev = n1
	list.Length++
}

func (list *DoublyLinkedList[T]) InsertBefore(n3 *DoublyLinkedListNode[T], value T) {
	newNode := &DoublyLinkedListNode[T]{
		Value: value,
	}

	if n3.Prev != nil {
		n3.Prev.Next = newNode
		newNode.Prev = n3.Prev
	} else {
		list.Head = newNode
	}

	n3.Prev = newNode
	newNode.Next = n3
	list.Length++
}
