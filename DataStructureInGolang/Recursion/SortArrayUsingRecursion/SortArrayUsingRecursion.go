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
	fmt.Println("Before Sorting")
	stk2.Display()
	fmt.Println("\nAfter Sorting")
	sortStack(stk2)
	stk2.Display()
}

func sortStack(stk2 *stack) {
	if stk2.Size == 1 {
		return
	}
	tmp := stk2.Top().(int)
	stk2.Pop()
	sortStack(stk2)
	insert(stk2, tmp)
}

func insert(stk2 *stack, tmp int) {
	if stk2.Size == 0 || stk2.Top().(int) <= tmp {
		stk2.Push(tmp)
		return
	}
	value := stk2.Top().(int)
	stk2.Pop()
	insert(stk2, tmp)
	stk2.Push(value)
	return
}

/*package main

func main() {
	userInput := make([]int, 0)
	//	userInput = []int{1, 3, 0, 5, -1}
	userInput = []int{2, 1}
	sortArrayUsingRecursion(userInput)
}

func sortArrayUsingRecursion(userInput []int) {
	if len(userInput) == 1 {
		return
	}
	tmp := userInput[len(userInput)-1]
	userInput = userInput[:len(userInput)-1] // Deleting last element.
	sortArrayUsingRecursion(userInput)
	insert(userInput, tmp)
	return
}

func insert(userInput []int, tmp int) {
	if len(userInput) == 0 || userInput[len(userInput)-1] <= tmp {
		userInput = append(userInput, tmp) // Appending element at the last
	}
	val := userInput[len(userInput)-1]
	userInput = userInput[:len(userInput)-1] // Deleting last element.
	insert(userInput, tmp)
	userInput = append(userInput, val)
	return
}
*/
/*package main

import (
	"container/vector"
)

var (
	vectorObject := vector.make(0)
	vevectorObject.
	listObject list.List = *list.New().Init()

)

func main() {
	arr := make([]int, 0)
	sortArray(arr)
}

func sortArray(arr []int) {
	//Base Condition: Already it is sorted.
	if len(arr) <= 1 {
		return
	}
	lastElement := arr[len(arr)-1]
	listObject.PushBack(lastElement)
	sortArray(arr)
	insert(arr, lastElement)
}

func insert(arr []int, element int) {
	if len(arr) == 0 || arr[len(arr)-1] <= element {
		listObject.PushBack(element)
		return
	}
	value := arr[len(arr)-1]
	arr = arr[0 : len(arr)-1] // Deleting last element
	insert(arr, value)
	arr = append(arr, value)
	return
}
*/
