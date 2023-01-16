package main

import "fmt"

type StackMethods interface {
	Push(value interface{})
	Peek() interface{}
	Pop()
	Size() int
	Reverse()
	Display()
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

func (s *Stack) Peek() interface{}{
	if s.size > 0 {
		fmt.Println("Element: ", s.node.value)
		return s.node.value
	} else {
		fmt.Println("No elements in stack to display.")
	}
	return nil
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

func (s *Stack) Reverse(){
	if s.node != nil{
		var next,prev *Node
		curr := s.node

		for curr != nil{
			next = curr.next
			curr.next = prev
			prev = curr
			curr = next
		}
		s.node = prev
		return
	}
}

func(s *Stack) Display(){
	if s.node != nil{
		ptr := s.node
		fmt.Println("Displaying elements")
		for ptr != nil{
			fmt.Print(ptr.value,"\n")
			fmt.Println("____")
			ptr= ptr.next
		}
		return
	}
	fmt.Println("Stack is empty.")
}

func main() {
	//Creating interface object.
	var stackObject StackMethods
	//Assigning Stack struct object to interface object.
	stackObject = new(Stack)

	stackObject.Push("A")
	fmt.Println("Size:", stackObject.Size())
	stackObject.Peek()
	stackObject.Peek()
	fmt.Println("Size:", stackObject.Size())
	stackObject.Push("B")
	stackObject.Push("C")
	fmt.Println("Size:", stackObject.Size())
	stackObject.Peek()
	fmt.Println("Size:", stackObject.Size())
	fmt.Println("Before Reversing:")
	stackObject.Display()
	fmt.Println("After Reversing:")
	stackObject.Reverse()
	stackObject.Display()
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
