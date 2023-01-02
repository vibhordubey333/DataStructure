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
	fmt.Println("Inserting element:", value)
	dataObject := Node{} //&Node{value: value, next: nil}
	dataObject.value = value
	//If nil insert at head
	if s.head == nil {
		fmt.Println("Inserting in front as list is nil")
		s.head = &dataObject
		s.size++
	} else {
		fmt.Println("Inserting in rear as one or more element is present")
		ptr := s.head
		for i := 0; i < s.size; i++ {
			if ptr.next == nil {
				ptr.next = &dataObject
				s.size++
				return
			}
			ptr = ptr.next
		}
	}

}

func (s *SinglyLinkedList) Search(value interface{}) {

}

func (s *SinglyLinkedList) Size() int {
	return s.size
}

func (s *SinglyLinkedList) Display() {
	if s.Size() > 0 {
		ptr := s.head
		for ptr != nil{
			fmt.Println("Element: ",ptr.value)
			ptr = ptr.next
		}
	} else {
		fmt.Println("LinkedList is empty.")
	}
}

func (s *SinglyLinkedList) Reverse() {

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
	llo.Insert("B")
	llo.Insert("C")
	llo.Insert("D")
	fmt.Println("SinglyLinkedList Size:", llo.Size())
	fmt.Println("Before Reversing: ")
	llo.Display()
	fmt.Println("After Reversing List:")
	llo.Reverse()
	llo.Display()
}
