package main

import (
	"container/list"
	"errors"
	"log"
	"sync"
)

type stack struct {
	lock     sync.RWMutex
	stackObj *list.List
}

//Wrapper onto the call of PushFront.
func (c *stack) Push(value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stackObj.PushFront(value)
}
func (c *stack) Size() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.stackObj.Len()
}

func (c *stack) Front() {
	defer panicHandler()
	if c.Size() > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		if val, ok := c.stackObj.Front().Value.(interface{}); ok {
			log.Println("From Front element is : ", val)
		}
	}
}

func (c *stack) Pop() error {
	if c.Size() > 0 {
		c.lock.Lock()
		defer c.lock.Unlock()
		ele := c.stackObj.Front()
		c.stackObj.Remove(ele)
		log.Println("Pop element: ", ele.Value)
		return nil
	} else {
		return errors.New("Stack is already empty")
	}
}

func (c *stack) Empty() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.stackObj.Len() == 0
}
func main() {
	stackObj := &stack{
		stackObj: list.New(),
	}
	log.Println("Push Operation.")
	stackObj.Push("A")
	stackObj.Push("N")
	stackObj.Push("!@##@")
	stackObj.Push("12321")
	stackObj.Push("12.321")
	log.Println("Stack size :", stackObj.Size())
	log.Println("Displaying stack elements. ")
	stackObj.Front()
	log.Println("Pop operation begins.")
	err := stackObj.Pop()
	if err != nil {
		log.Println("Error in pop operation : ", err)
	}

}
func panicHandler() {
	if r := recover(); r != nil {
		log.Println("Recovered: ", r)
	}
}
