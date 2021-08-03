package datas

import (
	"fmt"
)

const (
	_OUT_OF_RANGE = "runtime error: index out of range [%d] with length %d"
	_IS_EMPTY     = "runtime error: list is empty"
)

type synglyLikedNode struct {
	value interface{}
	next  *synglyLikedNode
}

type SynglyLinkedList struct {
	head *synglyLikedNode
	tail *synglyLikedNode
	size int
}

func (l *SynglyLinkedList) Slice() []interface{} {
	values := make([]interface{}, l.Size())
	node := l.head
	for i := 0; i < l.Size(); i++ {
		values[i] = node.value
		node = node.next
	}
	return values
}

func (l *SynglyLinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *SynglyLinkedList) Size() int {
	return l.size
}

func (l *SynglyLinkedList) tryIndex(index int) {
	if index >= l.Size() {
		panic(fmt.Sprintf(_OUT_OF_RANGE, index, l.Size()))
	}
}

func (l *SynglyLinkedList) getNode(index int) *synglyLikedNode {
	l.tryIndex(index)

	current := 0
	node := l.head
	for current < index {
		node = node.next
		current++
	}
	return node
}

func (l *SynglyLinkedList) Get(index int) interface{} {
	if index < 0 {
		index = l.Size() + index
	}
	return l.getNode(index).value
}

func (l *SynglyLinkedList) PushFront(value interface{}) {
	node := &synglyLikedNode{
		value: value,
		next:  l.head,
	}
	l.head = node
	if l.tail == nil {
		l.tail = node
	}
	l.size++
}

func (l *SynglyLinkedList) PopFront() interface{} {
	if l.Size() == 0 {
		panic(_IS_EMPTY)
	}

	first := l.getNode(0)
	second := first.next

	l.head = second
	if second == nil {
		l.tail = nil
	}

	l.size--
	return first.value
}

func (l *SynglyLinkedList) PushBack(value interface{}) {
	node := &synglyLikedNode{
		value: value,
		next:  nil,
	}
	if l.tail != nil {
		l.tail.next = node
	}
	l.tail = node
	if l.head == nil {
		l.head = node
	}
	l.size++
}

func (l *SynglyLinkedList) PopBack() interface{} {
	if l.Size() == 0 {
		panic(_IS_EMPTY)
	}
	last := l.tail
	var prev *synglyLikedNode = nil
	if l.Size() >= 2 {
		prev = l.getNode(-2)
	}

	l.tail = prev
	if prev == nil {
		l.head = nil
	}

	l.size--
	return last.value
}

func createSynglyLikedNodeSequence(values []interface{}) (start, end *synglyLikedNode) {
	if len(values) == 0 {
		return nil, nil
	}
	start = &synglyLikedNode{
		value: values[0],
		next:  nil,
	}
	end = start
	for _, value := range values[1:] {
		node := &synglyLikedNode{
			value: value,
			next:  nil,
		}
		end.next = node
		end = node
	}
	return
}

func (l *SynglyLinkedList) Splice(index int, delCount int, values ...interface{}) {
	if index < 0 || index+delCount > l.Size() {
		panic(fmt.Sprintf(_OUT_OF_RANGE, index, l.Size()))
	}

	before := l.head
	if index >= 1 {
		before = l.getNode(index - 1)
	}

	after := l.tail
	aindex := index + delCount
	if aindex < l.Size() {
		after = l.getNode(aindex)
	}
	start, end := createSynglyLikedNodeSequence(values)

	if len(values) == 0 {
		before.next = after
	} else {
		before.next = start
		end.next = after
	}
	l.size = l.Size() + len(values) - delCount
}
