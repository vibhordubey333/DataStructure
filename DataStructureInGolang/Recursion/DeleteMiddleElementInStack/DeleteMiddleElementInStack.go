package main

import (
	"fmt"
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
		fmt.Printf(" %v", tmp.data)
		tmp = tmp.next
	}
}

func New() *stack {
	stk := new(stack)
	stk.lock = &sync.Mutex{}

	return stk
}

func (stk *stack) Top() interface{} {
	if stk.head != nil {
		return stk.head.data
	}
	return stk.head
}

func main() {
	stk2 := New()
	arr := make([]int, 0)
	arr = []int{5, 2, 1, 0, -1, 7}
	fmt.Println("Array Elements: \n", arr)
	for i := len(arr) - 1; i >= 0; i-- {
		stk2.Push(arr[i])
	}
	fmt.Printf("Stack Elements:")
	stk2.Display()
	current := 0
	size := stk2.Size
	DeleteMiddleElementInStack(stk2, current, size)
	stk2.Display()

}

func DeleteMiddleElementInStack(stk *stack, current, size int) {
	if stk.Size == 0 || current == size {
		return //Base Condition
	}
	value := stk.Top()
	stk.Pop()
	DeleteMiddleElementInStack(stk, current+1, size) // Hypothesis

	if current != size/2 {
		fmt.Println("Current:", current, "Value:", value)
		stk.Push(value)
	}
	return
}
