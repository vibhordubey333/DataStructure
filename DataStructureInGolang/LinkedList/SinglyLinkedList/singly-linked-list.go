package main

import "fmt"

type SinglyLinkedListMethods interface {
	Insert(userInput interface{})
	Display()
	Size() int
	Remove(userInput interface{})
	Search(userInput interface{})
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
	payload := &Node{value: value, next: nil}

	if s.head == nil {
		s.head = payload
		s.size++
		fmt.Println("Inserted: ", value, " to empty linked list successfully.")
		return
	}

	ptr := s.head

	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = payload
	s.size++
	fmt.Println("Inserted:", value, " at the end of the linked list.")
}

func (s *SinglyLinkedList) Display() {
	ptr := s.head
	for ptr != nil {
		fmt.Print(ptr.value, " -> ")
		ptr = ptr.next
	}
	fmt.Print(" X ")
}

func (s *SinglyLinkedList) Search(userInput interface{}) {
	ptr := s.head

	if ptr != nil {
		for ptr != nil {
			if ptr.value == userInput {
				fmt.Println("Element found at: ", ptr.value)
				return
			}
			ptr = ptr.next
		}
	} else {
		fmt.Println("List is empty.")
		return
	}
	fmt.Println("Element not found")
}

func (s *SinglyLinkedList) Reverse() {
	if s.head != nil {
		var next, prev *Node
		curr := s.head

		for curr != nil {
			next = curr.next
			curr.next = prev
			prev = curr
			curr = next
		}
		s.head = prev
		s.Display()
		return
	}

	fmt.Println("No items present in list.")
}

func (s *SinglyLinkedList) Remove(userInput interface{}) {

	if s.head != nil {
		ptr := s.head
		for ptr.next != nil {
			if ptr.next.value == userInput {
				ptr.next = ptr.next.next
				fmt.Println("Element Removed")
				s.size--
				return
			}
			ptr = ptr.next
		}

	} else {
		fmt.Println("List is empty.")
	}

}

func (s *SinglyLinkedList) Size() int {
	return s.size
}

func main() {
	var interfaceObject SinglyLinkedListMethods
	interfaceObject = new(SinglyLinkedList)
	interfaceObject.Insert("A")
	interfaceObject.Insert("B")
	interfaceObject.Insert("C")
	fmt.Println("Size is :", interfaceObject.Size())
	interfaceObject.Display()
	fmt.Println()
	interfaceObject.Search("AN")
	interfaceObject.Reverse()
	interfaceObject.Remove("A")
	fmt.Println("After Removing element")
	interfaceObject.Display()
}
