package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Add(val T) {
	newL := List[T]{l.next, val}
	l.next = &newL
}

func (l *List[T]) MoveNext() {
	if l.next == nil {
		return
	}

	l.val = l.next.val
	l.next = l.next.next
}

func main() {
	l := List[string]{nil, "Hello"}
	l.Add("!")
	l.Add("World")
	l.Add(" ")

	for i := 0; i < 4; i++ {
		fmt.Print(l.val)
		l.MoveNext()
	}
}
