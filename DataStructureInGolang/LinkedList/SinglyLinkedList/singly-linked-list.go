package main

import "fmt"

type SinglyLinkedListMethods interface {
	Insert(interface{})
	Search(interface{})
	Size() int
	Display()
	Reverse()
}

type SinglyLinkedList struct {
	head *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func (s *SinglyLinkedList) Insert(value interface{}) {
	dataObject := &Node{value: value, next: nil}
	//If nil insert at head
	if s.head == nil {
		s.head = dataObject
	} else {
		ptr := s.head
		for ptr != nil {
			ptr = ptr.next
		}
		ptr = dataObject
	}
	s.size++
}

func (s *SinglyLinkedList) Search(value interface{}) {

}

func (s *SinglyLinkedList) Size() int {
	return s.size
}

func (s *SinglyLinkedList) Display() {
	fmt.Println("LinkedList size is : ", s.Size())
	if s.Size() > 0 {
		// Creating a temporary variable. To avoid manipulation.
		ptr := s
		for ptr.head != nil {
			fmt.Println("Elements: ", s.head.value)
			ptr.head = ptr.head.next
		}
	} else {
		fmt.Println("LinkedList is empty.")
	}
}

func (s *SinglyLinkedList) Reverse() {
	/*
		ptr := s.head
		var prev *Node

		for ptr != nil {
			ptr, prev, ptr.next = ptr.next, ptr, prev
		}
	*/
	curr := s.head
	var prev *Node
	var next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	s.head = prev
}

func main() {
	var llo SinglyLinkedListMethods
	llo = new(SinglyLinkedList)
	llo.Insert("A")
	llo.Display()
	llo.Insert("B")
	llo.Display()
	llo.Insert("C")
	llo.Display()
	fmt.Println("After Reversing List:")
	llo.Reverse()
	llo.Display()
}
