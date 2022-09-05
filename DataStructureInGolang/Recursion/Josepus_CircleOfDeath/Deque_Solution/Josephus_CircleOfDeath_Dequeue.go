package main

import (
	"fmt"
	"github.com/gammazero/deque"
)

//https://www.geeksforgeeks.org/josephus-problem-iterative-solution/
/*
The intuition of using the deque is quite obvious since we're traversing in a circular manner oveer the numbers. If we see a nimber which is modulus of k
then remove that number from operation.

Algorithm:

Push back all the numbers from the deque then start traversing from the front of the deque.
Keep a count of the numbers we encounter. Pop out the front element and if the count is not divisible by k then we've to
push the element again from the back of the queue else don't push back the element(this operation is similar to only deleting the element from the future operations).
Keep doing this until the deque size becomes one.

*/
var dqObject deque.Deque[interface{}]

func main() {
	n, k := 5, 2
	for i := 1; i <= n; i++ {
		dqObject.PushBack(int64(i))
	}
	fmt.Println("Person remaining at index : ", Josephus_Deque(n, k, dqObject))
}

func Josephus_Deque(n, k int, dqObject deque.Deque[interface{}]) int64 {
	count := 0
	for dqObject.Len() > 1 {
		currentPointer := (dqObject.Front()).(int64)
		dqObject.PopFront()
		count++

		// If the count is divisible by K then don't have to push back again
		if count%k == 0 {
			continue
		}
		//Else push back the element for the future operations
		dqObject.PushBack(currentPointer)

	}
	//Only element left in the queue will be returned.
	return dqObject.Back().(int64)
}
