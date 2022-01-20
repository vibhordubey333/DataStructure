package main

import "fmt"

/*
Output 13

*/
func main() {
	N := 14
	person, k, index := make([]int, 0), 14, 0
	k -= 1
	for i := 1; i <= N; i++ {
		person = append(person, i)
	}
	josephus(person, k, index)

}

func josephus(person []int, k, index int) {
	if len(person) == 1 {
		fmt.Println(person[0])
		return
	}
	index = ((index + k) % len(person))
	person = append(person[:index], person[index+1:]...)
	josephus(person, k, index)
	return
}
