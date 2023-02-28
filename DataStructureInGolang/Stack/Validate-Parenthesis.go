package main

import "fmt"

/*
(a + {b - [c + d] + e} - f * {g - h})}
*/
func main() {
	userInput := "{([{}])}" //"{(a + {b - ([c + d]) + e} - f * {g - h})}"
	fmt.Println(isExpressionValid(userInput))
}
func isExpressionValid(userInput string) bool {
	//expressionString := ""
	stackObject := new(Stack)
	for _, v := range userInput {
		if string(v) == "(" || string(v) == "{" || string(v) == "[" {
			//expressionString += string(v)
			stackObject.Push(string(v))
			continue
		}
		// Verify if stack is empty.
		if stackObject.size <= 0 {
			return false
		}

		switch string(v) {
		case ")":
			check := stackObject.Pop()
			if check == "{" || check == "[" {
				return false
				break
			}

		case "}":
			check := stackObject.Pop()
			if check == "(" || check == "[" {
				return false
				break
			}

		case "]":
			check := stackObject.Pop()
			if check == "(" || check == "{" {
				return false
				break
			}

		}
	}
	if stackObject.size > 0 {
		return false
	}
	return true
}

type Stack struct {
	head *Node
	size int
}
type Node struct {
	value interface{}
	next  *Node
}

func (s *Stack) Push(v interface{}) {
	s.head = &Node{
		value: v,
		next:  s.head,
	}
	s.size++
}
func (s *Stack) Pop() string {
	if s.size > 0 {
		tmp := s.head.value
		fmt.Println("Removing:", s.head.value)
		s.head = s.head.next
		s.size--
		return tmp.(string)

	} else {
		fmt.Println("No elements found.")
		return ""
	}
	return ""
}
