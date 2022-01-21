package main

import "fmt"

/*
Output 13
*/
func main() {
	/*
		N = Total no. of people.
		K = Person will be killed.
		index = Index where  the person will die.
	*/
	N := 5
	person, k, index := make([]int, 0), 2, 0
	k -= 1
	for i := 1; i <= N; i++ {
		person = append(person, i)
	}
	fmt.Println("\tPerson lists: ", person)
	josephus(person, k, index)

}

func josephus(person []int, k, index int) {
	if len(person) == 1 {
		fmt.Println("\tPerson Survived: ", person[0])
		return
	}
	index = ((index + k) % len(person))
	fmt.Print("\tKilling Person At Index: ", index, "    ")
	person = append(person[:index], person[index+1:]...)
	fmt.Println("People Remain Now:", person)
	josephus(person, k, index)
	return
}
