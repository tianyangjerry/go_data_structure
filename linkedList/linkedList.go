package linkedList

type linkedListNode[T comparable] struct {
	value T
	next  *linkedListNode[T]
}

type LinkedList[T comparable] struct {
	head *linkedListNode[T]
}

func (list *LinkedList[T]) Add(value T) {
	newNode := &linkedListNode[T]{value: value, next: nil}

	if list.head == nil {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
}

func (list *LinkedList[T]) Append(value T) {
	newNode := &linkedListNode[T]{value: value, next: nil}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head

		for current.next != nil {
			current = current.next
		}

		current.next = newNode
	}
}

func (list *LinkedList[T]) Contains(value T) bool {
	if list.head == nil {
		return false
	}

	if list.head.value == value {
		return true
	}

	current := list.head
	var previous *linkedListNode[T]

	for current != nil && current.value != value {
		previous = current
		current = current.next
	}

	if current == nil {
		return false
	}

	previous.next = current.next
	return true
}

func (list *LinkedList[T]) Delete(value T) bool {
	if list.head == nil {
		return false
	}

	if list.head.value == value {
		list.head = list.head.next
		return true
	}

	current := list.head
	var previous *linkedListNode[T]
	for current != nil && current.value != value {
		previous = current
		current = current.next
	}

	if current == nil {
		return false
	}

	previous.next = current.next
	return true
}

func (list *LinkedList[T]) Size() int {
	count := 0
	current := list.head

	for current.next != nil {
		current = current.next
		count++
	}

	return count
}
