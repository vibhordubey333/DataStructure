/*
Example 1:

Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:

Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/
//https://leetcode.com/problems/reverse-string/

package main

import "fmt"

func main() {
	input := []string{"h", "e", "l", "l", "o"}
	reverseString(input)
	inputString := "hello"
	revereStringCase(inputString)
}

func reverseString(s []string) {
	k := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[k-i]
		s[k-i] = temp
	}
	fmt.Println("Reversed String: ", s)
}

//For case when string is recieved as input.
func revereStringCase(inputString string) {
	s := []rune(inputString)
	k := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[k-i]
		s[k-i] = temp
	}
	fmt.Println("Result :", string(s))
}
