package main

import "fmt"

type Item struct {
	value interface{}
	next  *Item
}
type Stack struct {
	top  *Item
	size int
}

func (stackobj *Stack) IsEmpty() bool {
	if stackobj.size <= 0 {
		return true
	} else {
		return false
	}
}

func (stackobj *Stack) Len() int {
	return stackobj.size
}
func (stackobj *Stack) TopOfStack() {
	if stackobj.IsEmpty() != true {
		fmt.Println("Top Element: ", stackobj.top.value)
	} else {
		fmt.Println("Stack is empty.")
	}
}
func (stackobj *Stack) Pop() (value interface{}) {
	if stackobj.Len() > 0 {
		value = stackobj.top.value
		stackobj.top = stackobj.top.next
		stackobj.size--
		return
	}
	return nil
}
func (stackobj *Stack) Push(value interface{}) {
	stackobj.top = &Item{
		value: value,
		next:  stackobj.top,
	}
	stackobj.size++
}
func main() {
	stack := new(Stack)
	stack.Push("A")
	stack.Push(12)
	stack.TopOfStack()
	stack.Push(23.23)
	fmt.Println("IsEmpty:", stack.IsEmpty())
	for stack.Len() > 0 {
		fmt.Println(stack.Pop())
	}
}
