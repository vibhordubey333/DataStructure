package main

import (
	"fmt"
	"log"
	"sync"
)

func panicHandler() {
	if r := recover(); r != nil {
		fmt.Println("Recovered: ", r)
	}
}

type stack struct {
	lock     sync.RWMutex
	stackObj []interface{}
}

func (s *stack) Push(value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.stackObj = append(s.stackObj, value)
}
func (s *stack) Pop() {

	if s.Size() > 0 {
		s.lock.Lock()
		defer s.lock.Unlock()
		s.stackObj = s.stackObj[:s.Size()-1]
	} else {
		log.Println("Stack is already empty.")
	}
}
func (s *stack) Peep() {
	defer panicHandler()
	if s.Size() > 0 {
		s.lock.Lock()
		defer s.lock.Unlock()
		log.Println("Pop operation applied on element : ", s.stackObj[s.Size()-1])
	} else {
		log.Println("Stack is already empty.")
	}
}
func (s *stack) Size() int {
	defer panicHandler()
	return len(s.stackObj)
}
func main() {
	stackObj := &stack{
		stackObj: make([]interface{}, 0),
	}
	log.Println("Push operation begins.")
	//Push operation.
	stackObj.Push("XZY")
	stackObj.Push(123)
	stackObj.Push(123.123)
	stackObj.Peep()
	stackObj.Pop()
	stackObj.Peep()
	stackObj.Pop()
	stackObj.Pop()
	stackObj.Pop()
	stackObj.Peep()
}
