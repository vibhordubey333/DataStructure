// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type StackMethods interface {
	Push(value interface{})
	Peek()
	Pop()
	Size() int
}

type Stack struct {
	node *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
}

func (s *Stack) Size() int {
	return s.size
}
func (s *Stack) Push(value interface{}) {
	s.node = &Node{
		value: value,
		next:  s.node,
	}
	s.size++
}

func (s *Stack) Peek() {
	if s.size > 0 {
		fmt.Println("Element: ", s.node.value)
	} else {
		fmt.Println("No elements in stack to display.")
	}

}

func (s *Stack) Pop() {
	if s.size > 0 {
		fmt.Printf("Removing element %v from stack.\n", s.node.value)
		s.node = s.node.next
		s.size--
	} else {
		fmt.Println("No elements in stack to remove.")
	}
}

func main() {
	//Creating interface object.
	var stackObject StackMethods
	//Assigning Stack struct object to interface object.
	stackObject = new(Stack)

	stackObject.Push("A")
	fmt.Println("Size:", stackObject.Size())
	stackObject.Peek()
	stackObject.Pop()
	stackObject.Peek()
	fmt.Println("Size:", stackObject.Size())
	stackObject.Push("B")
	fmt.Println("Size:", stackObject.Size())
	stackObject.Peek()
}

/*
Output:

Size: 1
Element:  A
Removing element A from stack.
No elements in stack to display.
Size: 0
Size: 1
Element:  B
*/
