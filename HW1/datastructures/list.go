package list

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head   *Node
	Length int
}

func New() *LinkedList {
	var newList LinkedList
	newList.Length = 0
	return &newList
}

func (l *LinkedList) Append(val int) {
	var newNode Node
	newNode.Value = val
	if l.Head == nil {
		l.Head = &newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &newNode
	}
	l.Length++
}

func (l *LinkedList) Pop() {
	current := l.Head
	var prevcurrent *Node
	for current.Next != nil {
		prevcurrent = current
		current = current.Next
	}
	prevcurrent.Next = nil
	l.Length--
}

func (l *LinkedList) PrintVals() {
	var current Node
	current = *l.Head
	for i := 0; i <= (*l).Length; i++ {
		fmt.Println(current.Value)
		if current.Next == nil {
			break
		}
		current = *current.Next
	}
}

func (l *LinkedList) At(pos int) int { //нумерация позиций с 0
	current := l.Head
	for i := 0; i < pos; i++ {
		current = current.Next
	}
	return current.Value
}

func (l *LinkedList) Size() int {
	return l.Length
}

func (l *LinkedList) DeleteAt(pos int) {
	var prevnode *Node
	current := l.Head
	for i := 0; i < pos; i++ {
		prevnode = current
		current = current.Next
	}
	prevnode.Next = current.Next
	l.Length--
}

func (l *LinkedList) Update(pos, val int) {
	current := l.Head
	for i := 0; i < pos; i++ {
		current = current.Next
	}
	current.Value = val
}

func NewFromSlice(s []int) *LinkedList {
	listFromSlice := New()
	for i := 0; i < len(s); i++ {
		listFromSlice.Append(s[i])
	}
	return listFromSlice
}
