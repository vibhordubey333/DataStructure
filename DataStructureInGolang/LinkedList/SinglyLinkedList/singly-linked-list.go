// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type SinglyLinkedListMethods interface {
	Insert(interface{})
	Search(interface{})
	Size() int
	Display()
}

type SinglyLinkedList struct {
	node *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func (s *SinglyLinkedList) Insert(value interface{}) {
	//If nil insert at head
	if s.node == nil {
		s.node = &Node{
			value: value,
			next:  s.node,
		}
	}else{
		
	}
	s.size++
}

func (s *SinglyLinkedList) Search(value interface{}) {

}

func (s *SinglyLinkedList) Size() int {
	return s.size
}

func (s *SinglyLinkedList) Display() {
	if s.Size() > 0 {
		// Creating a temporary variable. To avoid manipulation.
		tmpObject := s
		for tmpObject.node != nil {
			fmt.Println("Elements: ", s.node.value)
			tmpObject.node = tmpObject.node.next
		}
	}
}

func main() {
	var llo SinglyLinkedListMethods
	llo = new(SinglyLinkedList)
	llo.Insert("A")
	llo.Display()
}
