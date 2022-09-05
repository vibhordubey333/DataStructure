package main

import "fmt"

/*
Output-
Person lists:  [1 2 3 4 5]
        Killing Person At Index: 1    People Remain Now: [1 3 4 5]
        Killing Person At Index: 2    People Remain Now: [1 3 5]
        Killing Person At Index: 0    People Remain Now: [3 5]
        Killing Person At Index: 1    People Remain Now: [3]
        Person Survived:  3
*/
/*
Output when k -= 1 operation is not performed we're just passing the value as it is.
Person lists:  [1 2 3 4 5]
        Killing Person At Index: 2    People Remain Now: [1 2 4 5]
        Killing Person At Index: 0    People Remain Now: [2 4 5]
        Killing Person At Index: 2    People Remain Now: [2 4]
        Killing Person At Index: 0    People Remain Now: [4]
        Person Survived:  4
*/

func main() {
	/*
		N = Total no. of people.
		K = Person will be killed.
		index = Index where  the person will die.
	*/
	N := 5
	person, k, index := make([]int, 0), 2, 0
	//Decrementing the position by 1. As we need to find out the person which has survived. If we won't do this operation we'll get the position of the person who was killed.
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
