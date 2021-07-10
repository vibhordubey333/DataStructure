// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
	)


func main() {

	userString := "RADAR"
	palindromeChecker(userString)
	optimizedpalindromeChecker(userString)
	userString = "MALAYALAM"
	palindromeChecker(userString)
	optimizedpalindromeChecker(userString)
	userString = "XMALAYALAM"
	palindromeChecker(userString)
	optimizedpalindromeChecker(userString)

}
func optimizedpalindromeChecker(userString string) {

	start := time.Now()

	j := len(userString) - 1
	flag := 0
	stringLength := len(userString)/2
	for i := 0; i < stringLength; i++ {
		if userString[i] == userString[j] {
			flag++
		}
		j--
	}

	if flag == stringLength {
		fmt.Println("Palindrome")
	} else {
		fmt.Println(" Not Palindrome")
	}
	elapsed := time.Since(start)
	fmt.Println("optimizedpalindromeChecker",elapsed)
}


func palindromeChecker(userString string) {
	start := time.Now()
	sliceArray := []rune(userString) // Cannot convert into []string array

	flag := 0
	sliceLength := len(userString) / 2
	for i := 0; i < sliceLength; i++ {
		if sliceArray[i] == sliceArray[(len(userString)-1)-i] {
			flag++
		}
	}

	if flag == sliceLength {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
	elapsed := time.Since(start)
	fmt.Println("palindromeChecker: ",elapsed)
}
