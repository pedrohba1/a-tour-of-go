package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

// Add appends a new value to the end of the list.
func (li *List[T]) Add(val T) {
	current := li
	for current.next != nil {
		current = current.next
	}

	current.next = &List[T]{val: val}
}

// String returns a string representation of the list.
func (li *List[T]) String() string {
	str := "["
	current := li
	for current != nil {
		if current != li {
			str += ", "
		}

		str += fmt.Sprintf("%v", current.val)
		current = current.next
	}

	str += "]"
	return str
}

func main() {
	li := &List[int]{
		val:  1,
		next: nil,
	}
	li.Add(2)
	li.Add(3)
	li.Add(4)
	fmt.Println(li)

}
