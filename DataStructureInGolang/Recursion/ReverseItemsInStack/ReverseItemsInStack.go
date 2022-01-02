package main

import (
	"fmt"
	"log"
	"sync"
)

type element struct {
	data interface{}
	next *element
}

type stack struct {
	lock *sync.Mutex
	head *element
	Size int
}

func (stk *stack) Push(data interface{}) {
	stk.lock.Lock()

	element := new(element)
	element.data = data
	temp := stk.head
	element.next = temp
	stk.head = element
	stk.Size++

	stk.lock.Unlock()
}

func (stk *stack) Pop() interface{} {
	if stk.head == nil {
		return nil
	}
	stk.lock.Lock()
	r := stk.head.data
	stk.head = stk.head.next
	stk.Size--

	stk.lock.Unlock()

	return r
}

func (stk *stack) Display() {
	tmp := stk.head
	for tmp != nil {
		fmt.Println(tmp.data)
		tmp = tmp.next
	}
}

func New() *stack {
	stk := new(stack)
	stk.lock = &sync.Mutex{}

	return stk
}

func main() {
	stk2 := New()
	stk2.Push(5)
	stk2.Push(4)
	stk2.Push(3)
	stk2.Push(2)
	stk2.Push(1)
	fmt.Println("Before Reverse:")
	stk2.Display()
	fmt.Println("After Reverse:")
	reverse(stk2)
	stk2.Display()
}

//Functionality to revese items of stack without another space.
func reverse(stk2 *stack) {
	if stk2.head == nil {
		log.Println("Stack is empty")
	}
	if stk2.Size == 1 {
		return
	}
	tmp := stk2.Pop()
	reverse(stk2)
	pushInStack(stk2, tmp)

}

func pushInStack(stk2 *stack, tmp interface{}) {
	if stk2.Size == 0 {
		stk2.Push(tmp)
		return
	}
	value := stk2.Pop()
	pushInStack(stk2, tmp)
	stk2.Push(value)
	return

}
