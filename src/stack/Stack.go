package stack

import "fmt"

type Stack[T any] struct {
	head *Node[T]
}
type Node[T any] struct {
	data     T
	nextItem *Node[T]
}

func (li *Stack[T]) Len() int {
	if li.head.nextItem == nil {
		return 1
	}
	item := li.head
	i := 1
	for item.nextItem != nil {
		item = item.nextItem
		i++
	}
	return i
}
func (li *Stack[T]) GetByIndex(index int) *Node[T] {
	if index >= li.Len() {
		panic("index out of bounds")
	}
	item := li.head
	for i := 1; i < index; i++ {
		item = item.nextItem
	}
	return item
}
func (li *Stack[T]) Push(value T) {
	list := &Node[T]{data: value, nextItem: nil}

	if li.head == nil {
		li.head = list
		return
	}

	item := li.head
	for item.nextItem != nil {
		item = item.nextItem
	}
	item.nextItem = list

}

func (li *Stack[T]) Pop() {

	SecondLast := li.GetByIndex(li.Len() - 1)
	SecondLast.nextItem = nil
}
func (l *Stack[T]) Show() {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.data)
		p = p.nextItem
	}
	fmt.Println("")
}
