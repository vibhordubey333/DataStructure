// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

/*
write a program to reverse words of given string

	i/p : Vibhor is a passionate software developer

	o/p	 developer software passionate a is Vibhor
*/
func main() {
	inputString := "Vibhor   is a passionate software developer"
	fmt.Println("reverseStringWithHandlingSpace:", reverseString(inputString))
	fmt.Println("ReverseString:", reverseString(inputString))
}

func reverseStringWithHandlingSpace(inputString string) string {
	stringSlice := strings.Split(inputString, " ")
	newString := ""
	for i := len(stringSlice) - 1; i >= 0; i-- {
		if stringSlice[i] == "" { //strings.Fields
			continue
		}
		newString += stringSlice[i] + " "
	}
	return newString
}

// Handling of multiple spaces using strings.Fields
func reverseString(inputString string) string {
	stringSlice := strings.Fields(inputString)
	newString := ""
	for i := len(stringSlice) - 1; i >= 0; i-- {
		newString += stringSlice[i] + " "
	}
	return newString
}

/*
Output:
reverseStringWithHandlingSpace: developer software passionate a is Vibhor 
ReverseString: developer software passionate a is Vibhor 
*/
